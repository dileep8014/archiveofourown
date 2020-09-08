package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog"
	"github.com/shyptr/archiveofourown/global"
	"github.com/uber/jaeger-client-go"
	"reflect"
)

type Model struct {
	Key    string
	Expire int
}

func New(t string, primary int64, expire ...int) Model {
	ex := 3600
	if len(expire) > 0 {
		ex = expire[0]
	}
	return Model{Key: fmt.Sprintf("%s:%d", t, primary), Expire: ex}
}

func (m Model) HGetAll(dst interface{}) error {
	conn := global.RedisPool.Get()
	defer conn.Close()
	values, err := redis.Values(conn.Do("HGETALL", m.Key))
	if err != nil {
		return err
	}
	if len(values) == 0 {
		return NilError
	}
	return redis.ScanStruct(values, dst)
}

func (m Model) HMSet(val interface{}) error {
	conn := global.RedisPool.Get()
	defer conn.Close()
	// 构造args
	args := []interface{}{m.Key}
	reflectVal := reflect.ValueOf(val)
	for i := 0; i < reflectVal.NumField(); i++ {
		field := reflectVal.Field(i)
		fieldType := reflect.TypeOf(val).Field(i)
		// if fieldType is ptr
		if fieldType.Type.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		name := fieldType.Name
		if tag, ok := fieldType.Tag.Lookup("redis"); ok {
			if tag == "-" {
				continue
			}
			name = tag
		}
		args = append(args, name, field.Interface())
	}
	reply, err := redis.String(conn.Do("HMSET", args...))
	if err != nil {
		return err
	}
	if reply != "OK" {
		return SetError
	}
	// 设置过期时间
	if m.Expire > 0 {
		conn.Do("EXPIRE", m.Key, m.Expire)
	}
	return nil
}

func (m Model) Del() error {
	conn := global.RedisPool.Get()
	defer conn.Close()

	affected, err := redis.Int(conn.Do("DEL", m.Key))
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("删除键值失败")
	}
	return nil
}

func (m Model) Cache(ctx context.Context, dst interface{}, call func() error) error {
	var err error
	parentSpan := opentracing.SpanFromContext(ctx)
	tr := parentSpan.Tracer()
	sp := tr.StartSpan("redis", opentracing.ChildOf(parentSpan.Context())).SetTag("redis.key", m.Key)
	spanContext := sp.Context().(jaeger.SpanContext)

	logger := global.Logger.With().Str("trace_id", spanContext.TraceID().String()).Logger().
		With().Str("span_id", spanContext.SpanID().String()).Logger().
		Hook(zerolog.HookFunc(func(_ *zerolog.Event, level zerolog.Level, message string) {
			ef := []log.Field{log.Event(level.String()), log.String("redis.op", message)}
			if level >= zerolog.ErrorLevel {
				ext.Error.Set(sp, true)
				ef = append(ef, log.Error(err))
			}
			sp.LogFields(ef...)
		}))

	err = m.HGetAll(dst)
	if err != nil {
		// 记录日志
		if err != NilError {
			logger.Error().Caller(1).Interface("key", m.Key).Err(err).Msg("HGetAll")
		}
		// 查库
		if err := call(); err != nil {
			return err
		}
		// 存入缓存
		err = m.HMSet(reflect.ValueOf(dst).Elem().Interface())
		if err != nil {
			logger.Error().Caller(1).Interface("key", m.Key).Err(err).Msg("HMSet")
		}
		return nil
	}
	return nil
}
