package middleware

import (
	"gin_mall/global"
	"github.com/golang-jwt/jwt"
	"time"
)

type claims struct {
	name string
	phone string
	jwt.StandardClaims
}


func CreateToken(name,phone string) (string,error) {
	claim := claims{
		name,
		phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second*3).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256,claim)
	return token.SignedString(global.PrivateKey)
}

func VerifyToken(tokenstring string) error {
	_,err := jwt.ParseWithClaims(tokenstring,&claims{}, func(token *jwt.Token) (interface{}, error) {
		return  global.PublicKey,nil
	})
	if err != nil {
		return err
	}
	return nil
}
