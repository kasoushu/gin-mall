package main

import (
	"gin_mall/global"
	"gin_mall/initialize"
	"github.com/gin-gonic/gin"
)





func main() {
	r := gin.Default()
	initialize.Initial()
	add :=global.CONFIG.Server.Address+":"+global.CONFIG.Server.Port
	r.Run(add)

}
