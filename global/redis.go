package global

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisPool *redis.Pool

func SetupRedisCache() {
	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisSetting.Host,
				redis.DialPassword(RedisSetting.Password),
				redis.DialConnectTimeout(time.Duration(RedisSetting.ConnectTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(RedisSetting.ReadTimeOut)*time.Second),
				redis.DialWriteTimeout(time.Duration(RedisSetting.WriteTimeOut)*time.Second),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     30,
		IdleTimeout: 240 * time.Second,
	}
}
