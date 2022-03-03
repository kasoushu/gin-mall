package initialize

import (
	"gin_mall/global"
	"github.com/go-redis/redis"
)

// redis 数据库初始化连接函数
func initClient() {
	//fmt.Println(add)
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     global.Config.RedisServer.Address,
		Password: global.Config.RedisServer.Password, // no password set
		DB:       global.Config.RedisServer.Database, // use default DB
	})
	_, err := global.RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	//_, err = global.RDB.LPush("kkk", struct {
	//	Name string
	//	Age  int
	//}{Name: "lll", Age: 11}).Result()
	//if err != nil {
	//	fmt.Println("push ", err)
	//}
	//re, err := global.RDB.Exists("lll").Result()
	//if err != nil {
	//	fmt.Println("mdb+ ", err)
	//}
	//fmt.Println("res", re)
}
