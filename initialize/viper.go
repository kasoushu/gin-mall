package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)
func InitViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if _,ok :=err.(viper.ConfigFileNotFoundError);ok{
		fmt.Println("Not found")
	}
	if err!=nil{
		fmt.Println(err)
		panic(err)
	}
	viper.SetDefault("server.port",8080)
	viper.SetDefault("server.address","127.0.0.1")
	viper.WatchConfig()
}
