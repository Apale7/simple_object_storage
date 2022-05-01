package main

import (
	"Apale7/simple_object_storage/handler"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/upload", handler.UniqueFile, handler.UploadFile)
		api.POST("/share", handler.ShareFile)
		api.GET("/download", handler.DownloadBySharing)
		api.GET("/file/list", handler.FileList)
	}
}
