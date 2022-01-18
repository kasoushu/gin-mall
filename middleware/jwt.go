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
	Id    uint64
	Phone string
	jwt.StandardClaims
}

func CreateToken(id uint64, phone string) (string, error) {
	claim := claims{
		id,
		phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}
	fmt.Println("create token id,phone", id, "   ", phone)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(global.PrivateKey)
}

func VerifyToken(tokenstring string) uint64 {
	tk, err := jwt.ParseWithClaims(tokenstring, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return global.PublicKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return 0
	}
	c := tk.Claims.(*claims)
	//fmt.Println("==========id================", c)
	return c.Id
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := c.Request.Header.Get("token")
		if tk == "" {
			model.Failed("token is empty", c)
			c.Abort()
			return
		}
		k := VerifyToken(tk)
		if k == 0 {
			model.Failed("token authorization error", c)
			c.Abort()
			return
		}
		c.Set("primary_id", k)
		//fmt.Println("===================================== ", k)
		c.Next()
	}
}
