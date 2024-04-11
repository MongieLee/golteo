package controller

import (
	"ginl/app/model"
	"ginl/common/dto"
	"ginl/db"
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
}

// GetUserById 根据Id获取用户信息
func (u *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user := &model.User{}
	tx := db.GormDb.Where("id = ?", id).Find(user)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	if tx.RowsAffected <= 0 {
		result.Failure(c, "用户不存在", gin.H{})
		return
	}
	result.Success(c, "查询成功", user)
}

// GetUsers 获取用户列表
func (u *UserController) GetUsers(c *gin.Context) {
	pagination := &dto.Pagination{}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	if err != nil {
		pageSize = 10
	}
	pagination.PageSize = pageSize
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	pagination.PageNum = pageNum
	var users []model.User
	db.GormDb.Limit(pagination.GetPageSize()).Offset(pagination.GetPageNum()).Find(&users)
	result.Success(c, "查询成功", users)
}

// ModifyUser 修改用户信息
func (u *UserController) ModifyUser(c *gin.Context) {
	mdyUser := &model.User{}
	err := c.ShouldBindJSON(mdyUser)
	if err != nil {
		result.Failure(c, err.Error(), gin.H{})
		return
	}
	tx := db.GormDb.Find(mdyUser)
	if tx.RowsAffected <= 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "用户不存在", gin.H{})
		return
	}
	tx = db.GormDb.Model(model.User{}).Where("id = ?", mdyUser.Id).Updates(&mdyUser)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	if tx.RowsAffected <= 0 {
		result.Failure(c, "修改失败", gin.H{})
		return
	}
	result.Success(c, "修改成功", gin.H{})
}

// DeleteUser 删除用户（软删除）
func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	tx := db.GormDb.Delete(&model.User{}, id)
	if tx.Error != nil {
		result.Failure(c, tx.Error.Error(), gin.H{})
		return
	}
	if tx.RowsAffected <= 0 {
		result.Failure(c, "用户【"+id+"】不存在", nil)
		return
	}
	result.Success(c, id+"删除成功", gin.H{})
}
