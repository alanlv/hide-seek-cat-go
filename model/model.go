package model

import (
	"sync"
	"time"
)

type UserInfo struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string `json:"user_name"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
