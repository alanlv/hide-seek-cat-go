package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

/**
中间件添加自定义处理逻辑1：
在请求和返回的header中插入x-request-id用于唯一标识一次HTTP请求
*/
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check for header, use it if exists.
		requestId := c.Request.Header.Get("X-Request-Id")
		// create request id with UUID4
		if requestId == "" {
			u4 := uuid.NewV4()
			requestId = u4.String()
		}
		// expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// set header.
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
