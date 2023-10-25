package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LevelUserAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		username := session.Get("Username")
		userRole := session.Get("UserRole")
		if userRole == nil && username == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}
		if userRole == "Customer" || userRole == "Merchant" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
