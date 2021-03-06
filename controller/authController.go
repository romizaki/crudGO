package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// authService service.AuthService
	// jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}
func (c *authController) Login(ctx *gin.Context) {
	fmt.Printf("masuk")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "masuk routing login",
	})
}
func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "masuk routing register",
	})
}
