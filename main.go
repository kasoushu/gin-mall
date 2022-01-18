package main

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/initialize"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type User struct {
	id       int
	name     string
	age      int
	phone    string
	password string
}

func main() {
	r := gin.Default()
	initialize.Initial(r)

	add := global.Config.Server.Address + ":" + global.Config.Server.Port
	defer global.MDB.Close()
	fmt.Println(viper.Get("newarrary.0.name"))
	r.Run(add)
}
