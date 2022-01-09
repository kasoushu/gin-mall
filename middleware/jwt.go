package middleware

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"github.com/gin-gonic/gin"
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
			ExpiresAt: time.Now().Add(time.Minute*60).Unix(),
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

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get("token")
		if tk=="" {
			model.Failed("token is empty",c)
			c.Abort()
			return
		}
		if err :=VerifyToken(tk);err!=nil {
			model.Failed("token authorization error",c)
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
