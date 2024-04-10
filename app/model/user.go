package model

import (
	"ginl/db"
	"gorm.io/gorm"
)

type User struct {
	Id                int64          `json:"id"`
	UserId            int64          `json:"userId"`
	UserName          string         `json:"userName" gorm:"column:username" validate:"required"`
	NickName          string         `json:"nickName" gorm:"column:nickname"`
	EncryptedPassword string         `json:"password" validate:"required"`
	Salt              string         `json:"salt"`
	Status            bool           `json:"status"`
	Avatar            string         `json:"avatar"`
	CreatedAt         db.LocalTime   `json:"createAt"`
	UpdatedAt         db.LocalTime   `json:"updateAt"`
	DeletedAt         gorm.DeletedAt `json:"-"`
}

// TableName gorm默认会用结构体名复数作为表名【users】，定义TableName方法可以自定义
func (u User) TableName() string {
	return "user"
}
