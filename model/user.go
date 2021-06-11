package model

import (
	"HideSeekCatGo/config"
	"HideSeekCatGo/utils"
	"fmt"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

//"github.com/go-playground/validator/v10"
type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"column:user_name;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// validate the fields
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = utils.Encrypt(u.Password)
	return err
}

// compare password
func (u *User) Compare(pwd string) (err error) {
	err = utils.Compare(u.Password, pwd)
	return err
}

// insert data to database.
func (u *User) Create() error {
	return DB.Self.Create(&u).Error
}

// delete user by user'id
func DeleteUser(id uint) error {
	user := User{}
	user.ID = id
	return DB.Self.Delete(&user).Error
}

// update user info.
func (u *User) Update() error {
	return DB.Self.Save(&u).Error
}

// list all users.
func ListUser(username string, offset, limit int) ([]*User, int64, error) {
	if limit == 0 {
		limit = config.DefaultLimit
	}
	users := make([]*User, 0)
	var count int64

	where := fmt.Sprintf("user_name like '%%%s%%'", username)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func GetUser(username string) (*User, error) {
	u := &User{}
	d := DB.Self.Where("user_name = ?", username).First(&u)
	return u, d.Error
}
