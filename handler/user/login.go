package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/model"
	"HideSeekCatGo/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登陆
// @Description 用户登陆，传入用户名和密码，成功则下发TOKEN。(JWT)
// @Produce  json
// @Param user_name body string true "UserName"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	/**
	1. 解析用户名和密码
	2. comapre对比加密的密码是否是数据库中存储的密码
	3. 相同则授权通过，签发token
	*/
	var u model.User
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	user, err := model.GetUser(u.UserName)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// compare password
	if err := utils.Compare(user.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// sign the jwt token.
	token, err := utils.Sign(c, utils.Context{
		ID:       user.ID,
		UserName: user.UserName,
	}, "")
	handler.SendResponse(c, nil, model.Token{
		Token: token,
	})

}
