package http

import (
	"backend-takehome/internal/common/log"
	commentUc "backend-takehome/internal/usecase/comment"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *HTTPHandler) CreateComment(c *gin.Context) {
	var request commentUc.CreateCommentRequest
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreateComment] Failed to read request body"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Decode the request body using json.Unmarshal
	if err := json.Unmarshal(body, &request); err != nil {
		log.LogError(errors.WithMessage(err, "[Handler][CreateComment] Failed to unmarshal JSON"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorName, err := GetUsernameFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.AuthorName = authorName

	err = h.CommentUsecase.CreateComment(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func (h *HTTPHandler) GetCommentByPostID(c *gin.Context) {
	postID := c.Param("id")
	postIDInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse post id"})
		return
	}

	resp, err := h.CommentUsecase.GetCommentByPostID(postIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment": resp})
}
