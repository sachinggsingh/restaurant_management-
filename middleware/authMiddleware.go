package middleware

import (
	"net/http"
	"resturnat-management/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing auth token",
			})
			ctx.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.First_Name)
		ctx.Set("last_name", claims.Last_Name)
		ctx.Set("uid", claims.Uid)

		ctx.Next()
	}
}
