package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romizaki/crudGO/config"
	"github.com/romizaki/crudGO/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDbConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helo",
		})
	})
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	r.POST("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "routing create order",
		})
	})
	r.GET("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "routing orders",
		})
	})
	r.GET("/order/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "routing orders by id",
		})
	})
	r.PUT("order/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "routing update by id",
		})
	})
	r.DELETE("order/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "routing delete by id",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
