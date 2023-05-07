package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/social-site-go/controllers"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/signup", controllers.SignUp())
		user.POST("/login", controllers.Login())
	}
}
