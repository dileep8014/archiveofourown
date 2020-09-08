package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/shyptr/archiveofourown/global"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"strconv"
	"time"
)

const user_key_segment = 1000

type UserCache struct {
	ID        int64     `redis:"id"`
	Username  string    `redis:"username"`
	Email     string    `redis:"email"`
	Password  string    `redis:"password"`
	Avatar    string    `redis:"avatar"`
	Gender    int       `redis:"gender"`
	Introduce string    `redis:"introduce"`
	Root      bool      `redis:"root"`
	WorksNums int64     `redis:"worksNums"`
	WorkDay   int64     `redis:"workDay"`
	Words     int64     `redis:"words"`
	FansNums  int64     `redis:"fansNums"`
	CreatedAt time.Time `redis:"createdAt"`
}

func (u *UserCache) Key() string {
	return fmt.Sprintf("user:%d", u.ID/user_key_segment)
}

func (u *UserCache) HMSetAll() (err error) {
	defer errwrap.Wrap(&err, "userCache.HMSetAll")

	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	reply, err := conn.Do("HMSET", u.Key(), "id:"+id, u.ID, "username:"+id, u.Username, "email:"+id, u.Email,
		"password:"+id, u.Password, "root:"+id, u.Root, "worksNums:"+id, u.WorksNums, "workDay:"+id, u.WorkDay, "words:"+id, u.Words,
		"fansNums:"+id, u.FansNums, "createdAt:"+id, u.CreatedAt.Format(time.RFC3339), "avatar:"+id, u.Avatar, "gender:"+id, u.Gender,
		"introduce:"+id, u.Introduce)
	s, err := redis.String(reply, err)
	if err != nil {
		return err
	}
	if s != "OK" {
		err = SetError
	}
	return
}

func (u *UserCache) HMSetField(fields map[string]interface{}) (err error) {
	defer errwrap.Wrap(&err, "userCache.HMSetField")

	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	args := []interface{}{u.Key()}
	for k, v := range fields {
		args = append(args, k+":"+id, v)
	}
	s, err := redis.String(conn.Do("HMSET", args...))
	if err != nil {
		return err
	}
	if s != "OK" {
		err = SetError
	}
	return
}

func (u *UserCache) HINCRBY(field string, increment int) (err error) {
	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	_, err = conn.Do("HINCRBY", u.Key(), field+":"+id, increment)
	return
}

func (u *UserCache) HGetAll() (err error) {
	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	values, err := redis.Values(conn.Do("HMGET", u.Key(), "id:"+id, "username:"+id, "email:"+id, "password:"+id,
		"root:"+id, "worksNums:"+id, "workDay:"+id, "words:"+id, "fansNums:"+id, "createdAt:"+id, "avatar:"+id,
		"gender:"+id, "introduce:"+id))
	if err != nil {
		return err
	}
	var createdAt string
	_, err = redis.Scan(values, &u.ID, &u.Username, &u.Email, &u.Password, &u.Root, &u.WorksNums, &u.WorkDay, &u.Words, &u.FansNums,
		&createdAt, &u.Avatar, &u.Gender, &u.Introduce)
	if err != nil {
		return
	}
	u.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	return
}

func (u *UserCache) DeleteField(fields ...string) (err error) {
	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	args := []interface{}{u.Key()}
	for _, f := range fields {
		args = append(args, f+":"+id)
	}
	_, err = conn.Do("HDEL", args...)
	return
}

func (u *UserCache) Delete() (err error) {
	conn := global.RedisPool.Get()
	defer conn.Close()
	id := strconv.Itoa(int(u.ID))
	_, err = conn.Do("HDEL", u.Key(), "id:"+id, "username:"+id, "email:"+id, "password:"+id, "root:"+id, "worksNums:"+id,
		"workDay:"+id, "words:"+id, "fansNums:"+id, "createdAt:"+id, "avatar:"+id, "gender:"+id, "introduce:"+id)
	return
}
