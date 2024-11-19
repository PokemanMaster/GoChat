package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORS  解决同源跨域的问题
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 允许来自 http://localhost:3000 的跨域请求
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 指定预检请求结果可以缓存的时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 允许的请求头
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// 允许携带身份凭证
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 允许的HTTP方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 对于预检请求，直接返回状态码 200
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
			return
		}

		// 否则，继续处理请求
		ctx.Next()
	}
}
