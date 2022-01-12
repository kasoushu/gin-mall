package main

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/initialize"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	if service.CreateCommodity(model.Product{
		CategoryId: 123,
		Price:      213,
		Amount:     123,
	}) {
		fmt.Println("successful!")
	}
	defer global.MDB.Close()
	r.Run(add)
}
