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

	user.GET("/say", func(c *gin.Context) {
		k, _ := c.Get("primary_id")
		c.JSON(200, gin.H{
			"msg": "welcome to my home",
			"id":  k,
		})
	})

	//管理员
	admin := g.Group("/admin")
	admin.Use(middleware.JwtAuth())

	admin.GET("/friend", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello",
			"id":  c.Param("id"),
		})
	})

	admin.GET("/get_product_info/:id", api.GetSingleProductInfo)
	admin.GET("/get_order_info/:id", api.GetSingleOrderInfo)
	admin.GET("/get_category_info/:id", api.GetCategoryInfo)
	admin.GET("/get_address_info/:id", api.GetSingleAddressInfo)

	admin.POST("/update_product/:id", api.UpdateProduct)
	admin.POST("/create_product", api.CreteProduct)
	admin.POST("/delete_product/:id", api.DeleteProduct)

	admin.POST("/update_category/:id", api.UpdateCategory)
	admin.POST("/create_category", api.CreteCategory)
	admin.POST("/delete_category/:id", api.DeleteCategory)

	admin.POST("/update_order/:id", api.UpdateOrder)
	admin.POST("/create_order", api.CreteOrder)
	admin.POST("/delete_order/:id", api.DeleteOrder)

	admin.POST("/update_address/:id", api.UpdateAddress)
	admin.POST("/create_address", api.CreteAddress)
	admin.POST("/delete_address/:id", api.DeleteProduct)

}
