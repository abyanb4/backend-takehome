package post

const (
	CreatePostQuery string = "INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)"
	UpdatePostQuery string = "UPDATE posts SET title = ?, content = ? WHERE id = ?"
	DeletePostQuery string = "DELETE FROM posts WHERE id = ?"

	GetPostByIDQuery       string = "SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE id = ?"
	GetPostByAuthorIDQuery string = "SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE author_id = ? ORDER BY created_at DESC"
	GetAllPostQuery        string = "SELECT id, title, content, author_id, created_at, updated_at FROM posts ORDER BY id"
)
