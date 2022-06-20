package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *grom.DB                  = config.SetupDatabaseConfig()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.closeDatabaseConfig(db)
	r := gin.Default()

	authRouter := r.Group("api/auth")
	{
		authRouter.POST("/login")
		authRouter.POST("/register")
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
