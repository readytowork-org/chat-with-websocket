package middlewares

import (
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func authJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID token not available"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(c, idToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("UUID", token.UID)
		c.Next()
	}
}
