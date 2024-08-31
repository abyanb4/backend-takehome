package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	jwtCommon "backend-takehome/internal/common/jwt"
)

var jwtSecret []byte

func init() {
	jwtSecretStr := os.Getenv("JWT_SECRET")
	jwtSecret = []byte(jwtSecretStr)
}

func tokenMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		claims, err := parseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token " + err.Error()})
			c.Abort()
			return
		}

		// Store user ID and username in context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		next(c)
	}
}

func parseToken(tokenStr string) (*jwtCommon.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtCommon.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwtCommon.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("parseToken - invalid token")
	}
	return claims, nil
}
