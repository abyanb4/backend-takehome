package post

import "time"

type CreatePostRequest struct {
	Title    string `json:"title" validate:"required,min=1,max=255"`
	Content  string `json:"content" validate:"required"`
	AuthorID string `json:"author_id" validate:""`
}

type UpdatePostRequest struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Content     string `json:"content" validate:"required"`
	RequestorID string `json:"requestor_id" validate:"required"`
}

type PostDetail struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Content     string          `json:"content"`
	AuthorID    string          `json:"author_id"`
	CommentList []CommentDetail `json:"comment_list"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type CommentDetail struct {
	ID         int64     `json:"id"`
	PostID     int64     `json:"post_id"`
	AuthorName string    `json:"author_name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
