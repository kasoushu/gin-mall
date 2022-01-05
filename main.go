package main

import (
	"gin_mall/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// redis 数据库初始化连接函数
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	r := gin.Default()
	config.SetRouters(r)
	r.Run()
}
