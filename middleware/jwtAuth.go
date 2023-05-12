package middleware

import (
	"diary_api/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
JWTAuthMiddleware function:

1. Executes helper.ValidateJWT function.

2. If helper.ValidateJWT function is successfully executed, the pending handlers is executed inside the calling handler.

FYI: https://gin-gonic.com/docs/examples/custom-middleware/
*/
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Executes helper.ValidateJWT function.
		err := helper.ValidateJWT(context)
		if err != nil {
			// If helper.ValidateJWT fails to execute, "Authentication required" is returned.
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Printf("##### ERROR #####")
			fmt.Println(err)
			// Prevents pending handlers from being called.
			context.Abort()
			return
		}
		// If helper.ValidateJWT function is successfully executed,
		// the pending handlers is executed inside the calling handler.
		context.Next()
	}
}
