package http

import (
	"backend-takehome/internal/common/log"
	postUc "backend-takehome/internal/usecase/post"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *HTTPHandler) CreatePost(c *gin.Context) {
	var request postUc.CreatePostRequest
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreatePost] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Decode the request body using json.Unmarshal
	if err := json.Unmarshal(body, &request); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreatePost] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.AuthorID = uuid

	err = h.PostUsecase.CreatePost(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func (h *HTTPHandler) GetPostByID(c *gin.Context) {
	postID := c.Param("id")

	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id must be integer"})
		return
	}

	resp, err := h.PostUsecase.GetPostByID(postIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": resp})
}

func (h *HTTPHandler) GetPostByAuthorID(c *gin.Context) {
	authorID := c.Param("id")

	resp, err := h.PostUsecase.GetPostByAuthorID(authorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": resp})
}

func (h *HTTPHandler) GetAllPost(c *gin.Context) {
	resp, err := h.PostUsecase.GetAllPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": resp})
}

func (h *HTTPHandler) UpdatePost(c *gin.Context) {
	var req postUc.UpdatePostRequest
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreatePost] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreatePost] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postID := c.Param("id")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id must be integer"})
		return
	}

	requestorID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.PostUsecase.UpdatePost(postUc.UpdatePostRequest{
		ID:          postIDInt,
		Title:       req.Title,
		Content:     req.Content,
		RequestorID: requestorID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (h *HTTPHandler) DeletePostByID(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id must be integer"})
		return
	}

	requestorID, err := GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.PostUsecase.DeletePostByID(requestorID, postIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": resp})
}
