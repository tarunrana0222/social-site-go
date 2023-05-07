package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/social-site-go/controllers"
	"github.com/tarunrana0222/social-site-go/middlewares"
)

func PostRoutes(router *gin.Engine) {
	post := router.Group("/post", middlewares.Authenticate())
	{
		post.POST("/create", controllers.CreatePost())
		post.DELETE("/delete/:id", controllers.DeletePost())
		post.GET("/get/:id", controllers.GetSinglePost())
		post.GET("/get/all", controllers.GetAllPost())
	}
}
