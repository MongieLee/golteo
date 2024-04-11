package dto

type RegisterDto struct {
	LoginDto
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type LoginDto struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
