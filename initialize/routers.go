package initialize

import (
	"gin_mall/api"
	"gin_mall/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	//user login,signup
	g.POST("/login", api.UserLogin)
	g.POST("/signup", api.UserSignUp)
	g.POST("/admin_login", api.AdminLogin)
	g.POST("/admin_signup", api.AdminSIgnUp)
	//user association
	user := g.Group("/:id")
	//
	user.Use(middleware.JwtAuth())
	//管理员
	admin := g.Group("/admin")
	admin.Use(middleware.JwtAuth())

	//user.GET("/", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"msg":"welcome to my home",
	//	})
	//})
	//user.GET("/friend", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"msg":"hello",
	//		"id":c.Param("id"),
	//	})
	//})

}
