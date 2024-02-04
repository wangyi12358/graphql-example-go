package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		fmt.Printf("token: %s\n", token)
		context.Next()
	}
}
