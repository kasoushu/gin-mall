package initialize

import (
	"gin_mall/global"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// redis 数据库初始化连接函数
func InitClient() {
	add := viper.GetString("redis_server.address")
	//fmt.Println(add)
	global.RDB = redis.NewClient(&redis.Options{
		Addr:add,
		Password: viper.GetString("redis_server.password"), // no password set
		DB:viper.GetInt("redis_server.database"),  // use default DB
	})
	_, err := global.RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
