package middlewares

import (
	"log"
	"mygram/auth"
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := auth.ValidateToken(c)
		log.Println(verifyToken)
		if err != nil {
			errors := helpers.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}

			response := helpers.APIResponse("Token invalid", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		c.Set("currentUser", verifyToken)
		c.Next()
	}
}
