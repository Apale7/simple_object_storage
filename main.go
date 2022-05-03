package main

import (
	"time"

	config "Apale7/simple_object_storage/config_loader"
	"Apale7/simple_object_storage/dal/mysql"
	"Apale7/simple_object_storage/dal/redis"
	"Apale7/simple_object_storage/handler"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var port string

func main() {
	Init()
	r := gin.Default()
	r.Use(cors.Default())
	defer r.Run(":" + port)
	r.LoadHTMLGlob("./static/*")

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

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"},
		MaxAge:          6 * time.Hour,
	}

	return cors.New(c)
}
