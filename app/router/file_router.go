package router

import (
	"ginl/app/controller"
	"github.com/gin-gonic/gin"
)

func InitFileRoute(e *gin.RouterGroup) {
	fileController := &controller.FileController{}
	{
		e.POST("/files", fileController.SingleFileUpload)
		e.POST("/multipleFiles", fileController.MultipleFileUpload)
	}
}
