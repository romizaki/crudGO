package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romizaki/crudGO/controller"
)

var (
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helo",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helo dari password",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}