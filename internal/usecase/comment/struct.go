package comment

type CreateCommentRequest struct {
	PostID     int64  `json:"post_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	AuthorName string `json:"author_id" validate:""`
}
