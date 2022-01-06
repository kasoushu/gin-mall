package initialize

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"gin_mall/global"
)

func Initial()  {
	initViper()
	initClient()
	initPrivateKey()
}
func initPrivateKey() {
	var err error
	global.PrivateKey,err= rsa.GenerateKey(rand.Reader,1024)
	if err!=nil {
		panic(err)
	}
	global.PublicKey = &global.PrivateKey.PublicKey
	fmt.Println(global.PrivateKey)
	fmt.Println(global.PublicKey)
}
