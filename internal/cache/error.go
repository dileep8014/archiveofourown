package cache

import "errors"

var (
	SetError = errors.New("设置键值失败")
	NilError = errors.New("查找的key不存在于HASH中")
)
