package post

import (
	commentDmn "backend-takehome/internal/domain/comment"
	postDmn "backend-takehome/internal/domain/post"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type PostUsecase struct {
	postDomain    *postDmn.PostDomain
	commentDomain *commentDmn.CommentDomain
}

func NewPostUsecase(postDmn *postDmn.PostDomain, commentDmn *commentDmn.CommentDomain) *PostUsecase {
	return &PostUsecase{postDomain: postDmn, commentDomain: commentDmn}
}
