package middleware

import (
	"crypto/rsa"
	"gin_mall/global"
)

func CreateToken() {

}

func getPrivateKey() {
	global.PrivateKey = rsa.GenerateKey()
}