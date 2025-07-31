package middleware

import (
	"dompetin-api/pkg/jwt"
	"dompetin-api/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	MsgAuthHeaderMissing = "authorization header missing"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.Respond(c, http.StatusUnauthorized, false, MsgAuthHeaderMissing, nil, nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateJWT(tokenStr)
		if err != nil {
			response.Respond(c, http.StatusUnauthorized, false, err.Error(), nil, nil)
			c.Abort()
			return
		}

		// Inject claims into context
		c.Set("user", claims)
		c.Next()
	}
}
