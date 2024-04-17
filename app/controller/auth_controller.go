package controller

import (
	"errors"
	"ginl/app/model"
	"ginl/app/model/dto"
	"ginl/config"
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct {
}

// Login 用户登陆
func (a *AuthController) Login(c *gin.Context) {
	var user model.User
	var inputDto dto.LoginDto
	err := c.ShouldBindJSON(&inputDto)
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		return
	}
	tx := config.Db.Where("username = ?", inputDto.UserName).Find(&user)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	if tx.RowsAffected == 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "账号或密码错误", gin.H{})
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(inputDto.Password+user.Salt)); err != nil {
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
	var inputModel dto.RegisterDto
	err := c.ShouldBindJSON(&inputModel)
	if err != nil {
		var vErr validator.ValidationErrors
		ok := errors.As(err, &vErr)
		if !ok {
			result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		}
		translate := vErr.Translate(utils.Trans)
		result.FailureWithCode(c, http.StatusBadRequest, "参数错误", utils.RemoveTopStruct(translate))
		return
	}
	tx := config.Db.Unscoped().Where("username = ?", inputModel.UserName).Find(&user)
	if tx.RowsAffected > 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "用户名已存在", gin.H{})
		return
	}
	salt, _ := utils.GenerateRandomSalt()
	password, err := utils.HashPassword(inputModel.Password, salt)
	if err != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	user.UserName = inputModel.UserName
	user.Salt = salt
	user.UserId = utils.GenerateSnowId()
	user.EncryptedPassword = password
	tx = config.Db.Create(&user)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	go func() {
		// 将用户列表全部缓存到redis中
		users := &[]model.User{}
		config.Db.Find(users)
		config.Rdb.Set("hot_users", users, 3600)
	}()
	result.SuccessWithData(c, gin.H{})
}

// RefreshToken 刷新用户Token
func (a *AuthController) RefreshToken(c *gin.Context) {
	var bodyJson map[string]interface{}
	err := c.ShouldBindJSON(&bodyJson)
	if err != nil {
		result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
		return
	}
	valid := false
	var refreshToken string
	if v, ok := bodyJson["refreshToken"]; ok {
		rv, ok := v.(string)
		if ok {
			refreshToken = rv
			valid = true
		}
	}
	if valid {
		userClaims, err := utils.ParseJWTToken(refreshToken)
		if err != nil {
			result.Failure(c, err.Error(), gin.H{})
			return
		}
		var user = &model.User{
			UserId:   userClaims.UserId,
			UserName: userClaims.Username,
		}
		newRefreshToken, err := utils.GenerateRefreshToken(user)
		if err != nil {
			result.Failure(c, err.Error(), gin.H{})
		}
		newAccessToken, err := utils.GenerateAccessToken(user)
		if err != nil {
			result.Failure(c, err.Error(), gin.H{})
		}
		result.SuccessWithData(c, model.Auth{
			Token:        newAccessToken,
			RefreshToken: newRefreshToken,
		})
	} else {
		result.FailureWithCode(c, http.StatusBadRequest, "参数异常", gin.H{})
	}
}
