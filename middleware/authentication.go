package middleware

import (
	"challenge_10/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)

		if err != nil {
			errorMessage := fmt.Sprintf("%s", err)
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
