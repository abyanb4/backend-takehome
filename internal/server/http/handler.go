package http

import (
	commentUc "backend-takehome/internal/usecase/comment"
	postUc "backend-takehome/internal/usecase/post"
	userUc "backend-takehome/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	PostUsecase    *postUc.PostUsecase
	UserUsecase    *userUc.UserUsecase
	CommentUsecase *commentUc.CommentUsecase
}

func RegisterHandler(r *gin.Engine, h *HTTPHandler) {

	// post
	r.GET("/posts/:id", tokenMiddleware(h.GetPostByID))
	r.GET("/posts/author/:id", tokenMiddleware(h.GetPostByAuthorID))
	r.GET("/posts", tokenMiddleware(h.GetAllPost))
	r.POST("/posts", tokenMiddleware(h.CreatePost))
	r.PUT("/posts/:id", tokenMiddleware(h.UpdatePost))
	r.DELETE("/posts/:id", tokenMiddleware(h.DeletePostByID))
	// comment
	r.POST("/posts/:id/comment", tokenMiddleware(h.CreateComment))
	r.GET("/posts/:id/comment", tokenMiddleware(h.GetCommentByPostID))

	// user
	r.POST("/user/register", h.RegisterUser)
	r.POST("/user/login", h.LoginUser)
	r.POST("/user/token", tokenMiddleware(h.DummyGetTokenUser))
}
