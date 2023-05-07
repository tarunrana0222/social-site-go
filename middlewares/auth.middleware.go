package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/social-site-go/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("x-token")

		if token == "" {
			ctx.JSON(401, gin.H{"message": "Auth headers missing"})
			ctx.Abort()
			return
		}

		userId, err := helpers.ValidateToken(token)
		if err != nil {
			ctx.JSON(401, gin.H{"message": "Auth Failed", "err": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
