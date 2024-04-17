package model

import (
	"gorm.io/gorm"
)

type User struct {
	Id                int64          `json:"id"`
	UserId            int64          `json:"userId"`
	UserName          string         `json:"username" gorm:"column:username" validate:"required"`
	NickName          string         `json:"nickname" gorm:"column:nickname"`
	EncryptedPassword string         `json:"-" validate:"required"`
	Salt              string         `json:"-"`
	Status            bool           `json:"status"`
	Avatar            string         `json:"avatar"`
	CreatedAt         LocalTime      `json:"createAt"`
	UpdatedAt         LocalTime      `json:"updateAt"`
	DeletedAt         gorm.DeletedAt `json:"-"`
}

// TableName gorm默认会用结构体名复数作为表名【users】，定义TableName方法可以自定义
func (u User) TableName() string {
	return "sys_user"
}
