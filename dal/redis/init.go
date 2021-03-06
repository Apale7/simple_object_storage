package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
)

func Init() {
	err := initClient()
	if err!= nil {
		panic(err)
	}
	logrus.Info("redis init success")
}

func initClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
