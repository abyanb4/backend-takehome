package comment

import "time"

type Comment struct {
	ID         int64     `json:"id" db:"id"`
	PostID     int64     `json:"post_id" db:"post_id"`
	AuthorName string    `json:"author_name" db:"author_name"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
