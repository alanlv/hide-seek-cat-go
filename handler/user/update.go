package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 更新用户信息
// @Description 传入用户的id和全部字段进行更新信息
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.User true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var u model.User
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u.ID = uint(userId)

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

	// save changed fields.
	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
