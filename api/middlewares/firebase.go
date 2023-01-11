package middlewares

import (
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FirebaseAuthMiddleWare struct {
	service services.FirebaseService
}

func NewFirebaseAuthMiddleware(
	service services.FirebaseService,

) FirebaseAuthMiddleWare {
	return FirebaseAuthMiddleWare{
		service: service,
	}
}

func (client FirebaseAuthMiddleWare) AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID token not available"})
			c.Abort()
			return
		}

		token, err := client.service.VerifyToken(idToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set(constants.UID, token.UID)
		c.Next()
	}
}
