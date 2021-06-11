package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/model"
	"HideSeekCatGo/utils"
	"github.com/gin-gonic/gin"
)

// login by jwt.
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
