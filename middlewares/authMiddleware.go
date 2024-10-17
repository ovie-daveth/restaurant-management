package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt" // Import the JWT package
)

var jwtSecret = []byte("674UYRBVF9487GVBKU-43987GYBFED-") // Replace with your actual secret key

// Authentication middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort() // Stop further processing
			return
		}

		// Validate the token (assuming it's in the format "Bearer <token>")
		if len(strings.Split(token, " ")) != 2 || strings.Split(token, " ")[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token = strings.Split(token, " ")[1] // Extract the token part

		// Parse the token
		claims := &jwt.StandardClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil // Return the secret key for validation
		})

		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Token is valid, proceed to the next handler
		c.Next()
	}
}
