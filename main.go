package main

import (
	"gin_mall/initialize"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	initialize.InitViper()
	initialize.InitClient()
	r.Run(viper.GetString("server.address")+":"+viper.GetString("server.port"))
}
