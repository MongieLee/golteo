package controller

import (
	"ginl/app/model"
	"ginl/common/dto"
	"ginl/config"
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
	tx := config.Db.Where("user_id = ?", id).Find(user)
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

// GetUsers 获取用户列表，存Redis
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
	if err = config.Rdb.Get("hot_users", &users); err != nil {
		tx := config.Db.Limit(pagination.GetPageSize()).Offset(pagination.GetPageNum() * pagination.GetPageSize()).Find(&users)
		if tx.Error != nil {
			result.Failure(c, "查询失败", gin.H{})
		} else {
			//go func() {
			//	// 将用户列表全部缓存到redis中
			//	users := &[]model.User{}
			//	config.Db.Find(users)
			//	config.Rdb.Set("hot_users", users, 3600)
			//}()
			result.Success(c, "查询成功", users)
		}
	} else {
		// 获取起始的索引
		start := pagination.GetPageSize() * pagination.GetPageNum()
		// 获取结束的索引，-1时因为索引从0开始
		end := start + pagination.GetPageSize() - 1
		// 如果结束的索引大于列表长度，则将索引设为最后一个
		if end >= len(users) {
			end = len(users) - 1
		}
		if start >= len(users) {
			result.Success(c, "查询成功", []model.User{})
			return
		}
		// 获取左包含右不包含的用户数据
		rUsers := users[start : end+1]
		result.Success(c, "查询成功", rUsers)
	}
}

// ModifyUser 修改用户信息
func (u *UserController) ModifyUser(c *gin.Context) {
	mdyUser := &model.User{}
	err := c.ShouldBindJSON(mdyUser)
	if err != nil {
		result.Failure(c, err.Error(), gin.H{})
		return
	}
	dbUser := &model.User{}
	tx := config.Db.Where("user_id = ?", mdyUser.UserId).Find(dbUser)
	if tx.RowsAffected <= 0 {
		result.FailureWithCode(c, http.StatusBadRequest, "用户不存在", gin.H{})
		return
	}
	tx = config.Db.Model(model.User{}).Where("user_id = ?", mdyUser.UserId).Updates(&mdyUser)
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
	tx := config.Db.Delete(&model.User{}, id)
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
