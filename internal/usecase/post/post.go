package post

import (
	postDmn "backend-takehome/internal/domain/post"
	"errors"
)

func (uc *PostUsecase) CreatePost(req CreatePostRequest) error {
	// Sanitize input
	err := ValidateCreatePostRequest(&req)
	if err != nil {
		return err
	}

	return uc.postDomain.CreatePostDB(&postDmn.Post{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorID,
	})
}

func (uc *PostUsecase) GetPostByID(id int64) (PostDetail, error) {
	postDetail, err := uc.postDomain.GetPostByIDDB(id)
	if err != nil {
		return PostDetail{}, err
	}

	postComment, err := uc.commentDomain.GetCommentByPostIDDB(postDetail.ID)
	if err != nil {
		return PostDetail{}, err
	}

	result := parsePostDetailResponse(postDetail, postComment)

	return result, nil
}

func (uc *PostUsecase) GetPostByAuthorID(authorID string) ([]postDmn.Post, error) {
	return uc.postDomain.GetPostByAuthorIDDB(authorID)
}

func (uc *PostUsecase) GetAllPost() ([]postDmn.Post, error) {
	return uc.postDomain.GetAllPostDB()
}

func (uc *PostUsecase) UpdatePost(req UpdatePostRequest) error {
	// check if the author and requestor has same id
	err := ValidateUpdatePostRequest(&req)
	if err != nil {
		return err
	}

	postDB, err := uc.postDomain.GetPostByIDDB(req.ID)
	if err != nil {
		return err
	}

	if postDB.ID == 0 {
		return errors.New("post not found")
	}

	if postDB.AuthorID != req.RequestorID {
		return errors.New("you are not authorized to update this post")
	}

	return uc.postDomain.UpdatePostDB(postDmn.Post{
		ID:      req.ID,
		Content: req.Content,
		Title:   req.Title,
	})
}

func (uc *PostUsecase) DeletePostByID(requestorID string, postID int64) (string, error) {
	// check if the author and requestor has same id
	postDB, err := uc.postDomain.GetPostByIDDB(postID)
	if err != nil {
		return "", err
	}

	if postDB.ID == 0 {
		return "", errors.New("post not found")
	}

	if requestorID != postDB.AuthorID {
		return "", errors.New("you are not authorized to delete this post")
	}

	return uc.postDomain.DeletePostDB(postID)
}
