package middleware

import (
	"diary_api/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ============================================================
// JWTAuthMiddleware function
// ============================================================
// FYI: https://gin-gonic.com/ja/docs/examples/custom-middleware/
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helper.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Println(err)
			context.Abort()
			return
		}
		context.Next()
	}
}
