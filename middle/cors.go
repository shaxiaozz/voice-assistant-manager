package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方法
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 添加跨域响应头
			c.Header("Content-Type", "application/json")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "X-Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, Cache-Control")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 放行OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
