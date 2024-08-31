package comment

import (
	commentDmn "backend-takehome/internal/domain/comment"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type CommentUsecase struct {
	commentDomain *commentDmn.CommentDomain
}

func NewCommentUsecase(repo *commentDmn.CommentDomain) *CommentUsecase {
	return &CommentUsecase{commentDomain: repo}
}
