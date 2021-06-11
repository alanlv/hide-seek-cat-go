package handler

import (
	"HideSeekCatGo/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
根据指定的返回格式，封装自己的返回函数
*/
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
