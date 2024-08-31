package http

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (string, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", errors.New("user id not found in token")
	}

	userIDStr := userID.(string)

	return userIDStr, nil
}

func GetUsernameFromContext(c *gin.Context) (string, error) {
	userID, exists := c.Get("username")
	if !exists {
		return "", errors.New("username not found in token")
	}

	userIDStr := userID.(string)

	return userIDStr, nil
}
