package model

import "time"

type Auth struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	ExpireAt     time.Time `json:"-"`
}
