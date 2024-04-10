package controller

import (
	"ginl/app/model"
	"ginl/db"
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

// Login 用户登陆
func (a *AuthController) Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		return
	}
	inputPassword := user.EncryptedPassword
	tx := db.GormDb.Where("username = ?", user.UserName).Find(&user)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	if tx.RowsAffected == 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "账号或密码错误", gin.H{})
		return
	}
	bcryptPassword, err := utils.HashPassword(inputPassword, user.Salt)
	if bcryptPassword != user.EncryptedPassword {
		result.FailureWithCode(c, http.StatusBadRequest, "账号或密码错误", gin.H{})
		return
	}
	token, err := utils.GenerateAccessToken(&user)
	if err != nil {
		result.FailureWithData(c, gin.H{})
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(&user)
	if err != nil {
		result.FailureWithData(c, gin.H{})
		return
	}
	result.SuccessWithData(c, &model.Auth{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

// Register 用户注册
func (a *AuthController) Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		return
	}
	err = utils.Validator.Struct(&user)
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		return
	}
	tx := db.GormDb.Unscoped().Where("username = ?", user.UserName).Find(&user)
	if tx.RowsAffected > 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "用户名已存在", gin.H{})
		return
	}
	salt, _ := utils.GenerateRandomSalt()
	password, err := utils.HashPassword(user.EncryptedPassword, salt)
	if err != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	user.Salt = salt
	user.UserId = utils.GenerateSnowId()
	user.EncryptedPassword = password
	tx = db.GormDb.Create(&user)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	result.SuccessWithData(c, gin.H{})
}

// RefreshToken 刷新用户Token
func (a *AuthController) RefreshToken(c *gin.Context) {
	result.SuccessWithData(c, gin.H{})
}
