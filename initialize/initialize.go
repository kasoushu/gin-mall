package initialize

import (
	"crypto/rand"
	"crypto/rsa"
	"gin_mall/global"
)

func Initial()  {
	initViper()
	initClient()
	initPrivateKey()
	initMysql()
}
func initPrivateKey() {
	var err error
	global.PrivateKey,err= rsa.GenerateKey(rand.Reader,1024)
	if err!=nil {
		panic(err)
	}
	global.PublicKey = &global.PrivateKey.PublicKey
}
