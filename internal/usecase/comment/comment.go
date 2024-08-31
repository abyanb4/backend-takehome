package comment

import commentDmn "backend-takehome/internal/domain/comment"

func (uc *CommentUsecase) CreateComment(req CreateCommentRequest) error {
	// Sanitize input
	err := ValidateCreatePostRequest(&req)
	if err != nil {
		return err
	}

	return uc.commentDomain.CreateCommentDB(commentDmn.Comment{
		PostID:     req.PostID,
		AuthorName: req.AuthorName,
		Content:    req.Content,
	})
}

func (uc *CommentUsecase) GetCommentByPostID(postID int64) ([]commentDmn.Comment, error) {
	return uc.commentDomain.GetCommentByPostIDDB(postID)
}
