package user

import (
	"HideSeekCatGo/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// create a new user.
func Create(c *gin.Context) {
	var r struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": errno.ErrBind,
		})
		return
	}
	log.Printf("username is: [%s], password is [%s]", r.UserName, r.Password)
	if r.UserName == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
	}
	if errno.IsErrUserNotFound(err) {
		log.Print("err type is ErrUserNotFound")
	}
	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
