package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
)

type Token struct{ V string }

func (t Token) SetEX(seconds int) (err error) {
	defer errwrap.Wrap(&err, "tokenCache.SetEX")

	conn := global.RedisPool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("SETEX", t.V, seconds, t.V))
	if err != nil {
		return
	}
	if s != "OK" {
		err = SetError
	}
	return
}

func (t Token) Exists() (bool, error) {
	conn := global.RedisPool.Get()
	defer conn.Close()
	n, err := redis.Int(conn.Do("EXISTS", t.V))
	if err != nil {
		return false, err
	}
	return n == 1, nil
}

func (t Token) Del() error {
	conn := global.RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", t.V)
	if err != nil {
		return err
	}
	return nil
}
