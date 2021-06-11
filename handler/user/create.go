package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/model"
	"github.com/gin-gonic/gin"
)

// @Summary 新建用户
// @Description 添加用户并保存到数据库中
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	/**
	1. 解析消息体
	2. 参数校验
	3. 加密密码
	4. 在数据库中添加数据记录
	5. 返回结果
	*/
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	// validate data.
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// insert the user to the database.
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	resp := CreateResponse{UserName: r.UserName}
	handler.SendResponse(c, nil, resp)

}
