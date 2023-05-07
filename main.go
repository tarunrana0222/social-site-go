package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/social-site-go/routes"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Static("/public", "./public")
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	routes.UserRoutes(router)
	routes.PostRoutes(router)
	router.Run("localhost:8080")
}
