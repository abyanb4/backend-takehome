package http

import (
	"backend-takehome/internal/common/log"
	userUc "backend-takehome/internal/usecase/user"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *HTTPHandler) RegisterUser(c *gin.Context) {
	var request userUc.RegisterUserRequest

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][RegisterUser] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Decode the request body using json.Unmarshal
	if err := json.Unmarshal(body, &request); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][RegisterUser] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UserUsecase.RegisterUser(userUc.RegisterUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][RegisterUser] Failed to register user"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_data": result})
}

func (h *HTTPHandler) LoginUser(c *gin.Context) {
	var request userUc.LoginRequest

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Decode the request body using json.Unmarshal
	if err := json.Unmarshal(body, &request); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UserUsecase.Login(userUc.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to register user"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_data": result})
}

func (h *HTTPHandler) DummyGetTokenUser(c *gin.Context) {
	var request userUc.LoginRequest

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Decode the request body using json.Unmarshal
	if err := json.Unmarshal(body, &request); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.UserUsecase.DummyGetToken(userUc.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][Login] Failed to register user"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_data": result})
}
