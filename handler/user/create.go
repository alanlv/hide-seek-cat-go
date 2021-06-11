package user

import (
	"HideSeekCatGo/errno"
	"HideSeekCatGo/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// create a new user.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	user := c.Param("username")
	log.Printf("URL username: %s", user)
	desc := c.Query("desc")
	log.Printf("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Printf("Header Content-Type: %s", contentType)

	log.Printf("username is: [%s], password is [%s]", r.UserName, r.Password)
	if r.UserName == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message."), nil)
		return
	}
	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}
	resp := CreateResponse{UserName: r.UserName}
	handler.SendResponse(c, nil, resp)
}
