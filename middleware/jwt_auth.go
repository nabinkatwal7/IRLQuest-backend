package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/helper"
)

func JWTAuthMiddleware() gin.HandlerFunc{
	return func(context *gin.Context){
		err := helper.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}