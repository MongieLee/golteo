package controller

import (
	"fmt"
	"ginl/service/result"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"time"
)

type FileController struct {
}

var storeDir = "./uploads"

// SingleFileUpload 单文件上传
func (f *FileController) SingleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.FailureWithData(c, gin.H{"msg": err.Error()})
		return
	}
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
}

// MultipleFileUpload 多文件上传
func (f *FileController) MultipleFileUpload(c *gin.Context) {
	form, multipartErr := c.MultipartForm()
	if multipartErr != nil {
		result.FailureWithData(c, gin.H{"msg": multipartErr.Error()})
		return
	}
	files := form.File["files"]
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
}

func generateFileName(prefix string, t time.Time) string {
	return fmt.Sprintf("%s%s%s", prefix, t.Format("20060102150405"), fmt.Sprintf("%d", t.Nanosecond()/1000))
}
