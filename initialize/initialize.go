package initialize

import (
	"crypto/rand"
	"crypto/rsa"
	"gin_mall/global"
	"github.com/gin-gonic/gin"
)

func Initial(g *gin.Engine)  {
	initViper()
	initClient()
	initPrivateKey()
	initMysql()
	initRouter(g)
}
func initPrivateKey() {
	var err error
	global.PrivateKey,err= rsa.GenerateKey(rand.Reader,1024)
	if err!=nil {
		panic(err)
	}
	global.PublicKey = &global.PrivateKey.PublicKey
}
