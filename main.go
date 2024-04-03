package main

import (
	"encoding/json"
	"fmt"
	"ginl/db"
	"ginl/entitys/persistent"
	"ginl/middleware"
	"ginl/model"
	"ginl/service/result"
	"ginl/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var storeDir = "./uploads"
var g errgroup.Group

func router02() http.Handler {
	e := gin.New()
	e.GET("/", func(c *gin.Context) {
		result.Success(c, "请求ok", gin.H{
			"fuck": "123",
		})
	})
	return e
}

func main() {
	dbErr := db.InitDb()
	if dbErr != nil {
		log.Println(dbErr.Error())
		return
	}
	r := gin.Default()
	r.Use(middleware.CorsHandler)
	//r.Use(middleware.CalcTimeHandler)
	//r.Use(middleware.GinBodyLogHandler)
	r.Static("/files", "./uploads")

	gGroup := r.Group("/api/v1")
	{
		authGroup := gGroup.Group("auth", func(c *gin.Context) {
			log.Println("auth组才有的中间件")
			c.Next()
		})
		authGroup.Use(func(c *gin.Context) {
			log.Println("auth组才有的中间件222")
			// 直接使用c是并发不安全的，通过c.Copy获取一份上下午文的副本
			contextCopy := c.Copy()
			go func() {
				log.Println("在goroutine中开启的输出打印" + contextCopy.Request.URL.Path)
			}()
			c.Next()
		})
		authGroup.POST("/login", func(c *gin.Context) {
			var user persistent.User
			err := c.ShouldBindJSON(&user)
			if err != nil {
				result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
				return
			}
			if user.UserName != "jack" || user.Id != 123 {
				result.FailureWithCode(c, http.StatusBadRequest, "账号或密码错误", gin.H{})
				return
			}
			token, err := utils.GenRegisteredClaims(&user)
			if err != nil {
				result.FailureWithData(c, gin.H{})
			}
			result.SuccessWithData(c, &model.AuthDto{
				Token: token,
			})
		})
		authGroup.POST("/register", func(c *gin.Context) {
			c.Next()
		}, func(c *gin.Context) {
			var user persistent.User
			err := c.ShouldBindJSON(&user)
			if err != nil {
				result.FailureWithCode(c, http.StatusBadRequest, err.Error(), gin.H{})
				return
			}
			sqlStr := "insert into user(username,encrypted_password) values (?,?)"
			exec, err := db.Db.Exec(sqlStr, user.UserName, user.EncryptedPassword)
			if err != nil {
				result.Failure(c, err.Error(), gin.H{})
				return
			}
			theId, err := exec.LastInsertId()
			if err != nil {
				result.FailureWithData(c, gin.H{})
				return
			}
			log.Println("theId is %v", theId)
			result.SuccessWithData(c, gin.H{})
		})
		authGroup.POST("/refreshToken", func(c *gin.Context) {
			result.SuccessWithData(c, gin.H{})
		})

		gGroup.GET("/getUser", func(c *gin.Context) {
			querySql := "select * from user"
			var user persistent.User
			var statusRaw []byte
			var createAt string
			var updateAt string
			err := db.Db.QueryRow(querySql).Scan(&user.Id, &user.UserName, &user.NickName, &user.EncryptedPassword, &user.Avatar, &statusRaw, &createAt, &updateAt)
			if err != nil {
				result.Failure(c, err.Error(), gin.H{})
				return
			}
			result.Success(c, "查询成功", user)
		})

		gGroup.GET("/getUsers", func(c *gin.Context) {
			querySql := "select * from user"
			var statusRaw []byte
			var createAt string
			var updateAt string
			rows, err := db.Db.Query(querySql)
			if err != nil {
				return
			}
			var users []persistent.User
			for rows.Next() {
				var user persistent.User
				err := rows.Scan(&user.Id, &user.UserName, &user.NickName, &user.EncryptedPassword, &user.Avatar, &statusRaw, &createAt, &updateAt)
				if err != nil {
					result.Failure(c, err.Error(), gin.H{})
					return
				}
				users = append(users, user)
			}
			result.Success(c, "查询成功", users)
		})

		gGroup.PATCH("/updateUser", func(c *gin.Context) {
			var body map[string]interface{}
			err := json.NewDecoder(c.Request.Body).Decode(&body)
			if err != nil {
				result.Failure(c, err.Error(), gin.H{})
				return
			}
			querySql := "update user set nickname = ? where id > ?"
			ret, err := db.Db.Exec(querySql, body["username"], body["id"])
			if err != nil {
				result.Failure(c, err.Error(), gin.H{})
				return
			}
			affected, err := ret.RowsAffected()
			if err != nil {
				return
			}
			msg := fmt.Sprintf("影响的行数:%d", affected)
			result.Success(c, "修改成功"+msg, gin.H{})
		})

		gGroup.DELETE("/users/:id", func(c *gin.Context) {
			id := c.Param("id")
			querySql := "delete from user where id = ?"
			ret, err := db.Db.Exec(querySql, id)
			if err != nil {
				result.Failure(c, err.Error(), gin.H{})
				return
			}
			affected, err := ret.RowsAffected()
			if err != nil {
				return
			}
			if affected > 0 {
				msg := fmt.Sprintf("影响的行数:%d", affected)
				result.Success(c, "删除成功"+msg, gin.H{})
			} else {
				result.FailureWithCode(c, http.StatusBadRequest, "删除失败，id不存在", gin.H{})
			}

		})

	}
	r.GET("/haha", middleware.AuthHandler, func(c *gin.Context) {
		c.Request.URL.Path = "/hi"
		r.HandleContext(c)
	})

	r.Any("/testAny", func(c *gin.Context) {
		method := c.Request.Method
		switch method {
		case "POST":
			log.Println("post")
		case "GET":
			log.Println("get")
		case "PUT":
			log.Println("put")
		case "HEAD":
			log.Println("head")
		case "OPTIONS":
			log.Println("options")
		default:
			log.Println("default")
		}
		result.SuccessWithData(c, gin.H{})
	})

	r.GET("/hi", func(c *gin.Context) {
		user := &persistent.User{}
		user.UserName = "aaa"
		user.EncryptedPassword = "123"
		user.Id = 111
		user.CreateAt = time.Now()
		user.UpdateAt = time.Now()

		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"code": 200,
			"data": user,
		})
	})

	r.GET("/getByte", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{Label: &label, Reps: reps}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})

	r.GET("/someXML", func(c *gin.Context) {
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})

	r.POST("/getPath/:username/:id", func(c *gin.Context) {
		username := c.Param("username")
		id := c.Param("id")
		log.Println("username:" + username + ",id:" + id)
		result.SuccessWithData(c, gin.H{})

	})

	r.POST("/uploadFile", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			result.FailureWithData(c, gin.H{"msg": err.Error()})
			return
		}
		storeDir := "./uploads"
		osErr := os.MkdirAll(storeDir, os.ModePerm)
		if osErr != nil {
			result.FailureWithData(c, gin.H{"msg": osErr.Error()})
			return
		}
		targetFileName := generateFileName("", time.Now()) + filepath.Ext(file.Filename)

		filePath := filepath.Join(storeDir, targetFileName)
		fileErr := c.SaveUploadedFile(file, filePath)
		if fileErr != nil {
			result.FailureWithData(c, gin.H{"msg": fileErr.Error()})
			return
		}
		result.Success(c, "文件上传成功", gin.H{
			"filePath": "http://localhost:8080/files/" + targetFileName,
		})
	})

	r.POST("/multipleUploadFile", func(c *gin.Context) {
		form, multipartErr := c.MultipartForm()
		if multipartErr != nil {
			result.FailureWithData(c, gin.H{"msg": multipartErr.Error()})
			return
		}
		files := form.File["file"]
		osErr := os.MkdirAll(storeDir, os.ModePerm)
		if osErr != nil {
			result.FailureWithData(c, gin.H{"msg": osErr.Error()})
			return
		}
		var fileNames []string
		for _, file := range files {
			targetFileName := generateFileName("", time.Now()) + filepath.Ext(file.Filename)
			fileNames = append(fileNames, "http://localhost:8080/files/"+targetFileName)
			filePath := filepath.Join(storeDir, targetFileName)
			fileErr := c.SaveUploadedFile(file, filePath)
			if fileErr != nil {
				result.FailureWithData(c, gin.H{"msg": fileErr.Error()})
				return
			}
		}
		result.Success(c, "文件上传成功", gin.H{
			"files": fileNames[:2],
		})
	})

	r.GET("/testRedirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})

	// 实现效果一致的
	r.GET("/customRedirect", func(c *gin.Context) {
		c.Header("Location", "https://baidu.com")
		c.Status(http.StatusMovedPermanently)
	})

	server2 := &http.Server{
		Addr:         ":8888",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server2.ListenAndServe()
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func generateFileName(prefix string, t time.Time) string {
	a := fmt.Sprintf("%03d", t.Nanosecond())
	fmt.Sprintf(a)
	return fmt.Sprintf("%s%s%s", prefix, t.Format("20060102150405"), fmt.Sprintf("%d", t.Nanosecond()/1000))
}
