package middleware

import (
	"net/http"
	"strings"

	"github.com/azevedoguigo/projeto-odonto-back.git/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required!",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
