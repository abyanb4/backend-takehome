package comment

import "github.com/jmoiron/sqlx"

type CommentDomain struct {
	DB *sqlx.DB
}

func NewCommentDomain(db *sqlx.DB) *CommentDomain {
	return &CommentDomain{DB: db}
}

func (d *CommentDomain) CreateCommentDB(comment Comment) error {
	stmt, err := d.DB.Prepare(CreateCommentQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.PostID, comment.AuthorName, comment.Content)
	if err != nil {
		return err
	}

	return err
}

func (d *CommentDomain) GetCommentByPostIDDB(postID int64) (result []Comment, err error) {
	err = d.DB.Select(&result, GetCommentByPostIDQuery, postID)
	if err != nil {
		return result, err
	}

	return result, nil
}
