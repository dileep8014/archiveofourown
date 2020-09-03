package v1

import "github.com/gin-gonic/gin"

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) Router(api gin.IRouter) {
	// 获取当前登录用户信息
	api.GET("/user/currentUser")
	// 获取指定用户信息
	api.GET("/user/:id")
	//
}
