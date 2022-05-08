package main

import (
	"Apale7/simple_object_storage/handler"
	"Apale7/simple_object_storage/middleware"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/login", handler.Login)
			user.POST("/register", handler.Register)
		}
		api.POST("/upload", middleware.JWTAuthMiddleware, handler.UniqueFile, handler.UploadFile)
		api.POST("/share", middleware.JWTAuthMiddleware, handler.ShareFile)
		api.GET("/download/share", handler.DownloadBySharing)
		api.GET("/download/self", middleware.JWTAuthMiddleware, handler.FileExist, handler.DownloadSelf)
		api.GET("/file/list", middleware.JWTAuthMiddleware, handler.FileList)
	}
	r.GET("/list", handler.ListPage)
}
