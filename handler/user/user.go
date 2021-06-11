package user

import (
	"HideSeekCatGo/model"
)

type CreateRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateResponse struct {
	UserName string `json:"user_name"`
}

type ListRequest struct {
	UserName string `json:"user_name"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
