package v1

import "github.com/gin-gonic/gin"

type User struct {
}

func NewUser() User {
	return User{}
}

func (u User) Router(api *gin.RouterGroup) {

}
