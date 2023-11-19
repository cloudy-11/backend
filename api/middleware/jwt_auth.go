package middleware

import (
	"github.com/cloudy-11/backend/domain"
	"github.com/cloudy-11/backend/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	X_USER_ID   = "x-user-id"
	X_USER_ROLE = "x-user-role"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := internal.IsAuthorized(authToken, secret)
			if authorized {
				role, userID, err := internal.ExtractIDAndRoleFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set(X_USER_ID, userID)
				c.Set(X_USER_ROLE, role)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
