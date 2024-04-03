package model

import "time"

type AuthDto struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	ExpireAt     time.Time `json:"-"`
}
