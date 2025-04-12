package middleware

import (
	"github.com/gin-gonic/gin"
)

// NewCORS 创建一个CORS跨域中间件，支持动态配置
func NewCORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 设置CORS头部
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
