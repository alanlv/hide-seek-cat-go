package middleware

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/utils"
	"github.com/gin-gonic/gin"
)

/**
jwt middleware.
*/
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := utils.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
