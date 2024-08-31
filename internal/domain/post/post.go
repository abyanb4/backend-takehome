package post

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostDomain struct {
	DB *sqlx.DB
}

func NewPostDomain(db *sqlx.DB) *PostDomain {
	return &PostDomain{DB: db}
}

func (d *PostDomain) CreatePostDB(post *Post) error {
	stmt, err := d.DB.Prepare(CreatePostQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return err
	}

	return err
}

// this should return the comment also
func (d *PostDomain) GetPostByIDDB(id int64) (result Post, err error) {
	var post []Post

	err = d.DB.Select(&post, GetPostByIDQuery, id)
	if err != nil {
		return result, err
	}

	if len(post) > 0 {
		result = post[0]
	}

	return result, nil
}

func (d *PostDomain) GetPostByAuthorIDDB(authorID string) (result []Post, err error) {
	err = d.DB.Select(&result, GetPostByAuthorIDQuery, authorID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (d *PostDomain) GetAllPostDB() (result []Post, err error) {
	err = d.DB.Select(&result, GetAllPostQuery)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (d *PostDomain) UpdatePostDB(post Post) error {
	stmt, err := d.DB.Prepare(UpdatePostQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *PostDomain) DeletePostDB(id int64) (string, error) {
	stmt, err := d.DB.Prepare(DeletePostQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Rows Affected = %d", rowsAffected), nil
}
