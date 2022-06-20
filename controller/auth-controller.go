package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.statusOK, gin.H{
		"message": "Hello ini login",
	})
}

func (c *authController) Register(ctx *gin.Context) {

}
