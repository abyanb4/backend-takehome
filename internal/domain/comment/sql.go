package comment

const (
	CreateCommentQuery string = "INSERT INTO comments (post_id, author_name, content) VALUES (?, ?, ?)"
	DeleteCommentQuery string = "DELETE FROM comments WHERE id = ?"

	GetCommentByPostIDQuery string = "SELECT id, post_id, author_name, content, created_at, updated_at FROM comments WHERE post_id = ? ORDER BY created_at DESC"
)
