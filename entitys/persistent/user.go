package persistent

import (
	"ginl/db"
)

type User struct {
	Id                int64         `json:"id"`
	UserName          string        `json:"userName" validate:"required"`
	NickName          string        `json:"nickName"`
	EncryptedPassword string        `json:"encryptedPassword" validate:"required"`
	Status            bool          `json:"status"`
	Avatar            string        `json:"avatar"`
	CreatedAt         db.LocalTime  `json:"createAt"`
	UpdatedAt         *db.LocalTime `json:"updateAt"`
}

// TableName gorm默认会用结构体名复数作为表名【users】，定义TableName方法可以自定义
func (u User) TableName() string {
	return "user"
}
