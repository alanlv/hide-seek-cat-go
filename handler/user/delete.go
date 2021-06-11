package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"HideSeekCatGo/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 删除某个用户
// @Description 拿到要删除的用户的id,然后根据id删除该用户。注意，gorm的软删除
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
