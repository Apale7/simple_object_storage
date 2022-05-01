package main

import (
	config "Apale7/simple_object_storage/config_loader"
	"Apale7/simple_object_storage/dal/mysql"
	"Apale7/simple_object_storage/dal/redis"
	"Apale7/simple_object_storage/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var port string

func main() {
	Init()
	r := gin.Default()
	defer r.Run(":" + port)

	register(r)
}

func Init() {
	logrus.SetReportCaller(true)
	config.Init()
	handler.Init()
	port = config.Get("port")
	if port == "" {
		port = "6789"
	}
	redis.Init()
	mysql.Init()
}
