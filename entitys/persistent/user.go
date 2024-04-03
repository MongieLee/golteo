package persistent

import "time"

type User struct {
	Id                int64     `json:"id"`
	UserName          string    `json:"userName" validate:"required"`
	NickName          string    `json:"nickName"`
	EncryptedPassword string    `json:"encryptedPassword" validate:"required"`
	Status            bool      `json:"status"`
	Avatar            string    `json:"avatar"`
	CreateAt          time.Time `json:"createAt"`
	UpdateAt          time.Time `json:"updateAt"`
}
