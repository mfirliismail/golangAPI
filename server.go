package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfirliismail/golangAPI/config"
	"github.com/mfirliismail/golangAPI/controller"
	"gorm.io/gorm"
)

//import db dan controller
var (
	db             *gorm.DB                  = config.SetupDatabaseConfig()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConfig(db)
	r := gin.Default()

	// (router ) bikin router berdasarkan controller
	authRouter := r.Group("api/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
