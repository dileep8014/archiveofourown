package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8000") // 设置允许访问所有域
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		ctx.Header("Access-Control-Expose-Headers", "*")
		ctx.Header("Access-Control-Max-Age", "172800")
		ctx.Header("Access-Control-Allow-Credentials", "false")
		if ctx.Request.Method == http.MethodOptions {
			ctx.JSON(http.StatusOK, gin.H{})
			return
		}
		ctx.Next()
	}
}
