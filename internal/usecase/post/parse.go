package post

import (
	commentDmn "backend-takehome/internal/domain/comment"
	postDmn "backend-takehome/internal/domain/post"
)

func parsePostDetailResponse(post postDmn.Post, comments []commentDmn.Comment) PostDetail {
	var result PostDetail
	var commentDetailList []CommentDetail

	result.ID = post.ID
	result.Title = post.Title
	result.Content = post.Content
	result.AuthorID = post.AuthorID
	result.CreatedAt = post.CreatedAt
	result.UpdatedAt = post.UpdatedAt

	for _, c := range comments {
		commentDetailList = append(commentDetailList, CommentDetail{
			ID:         c.ID,
			PostID:     c.PostID,
			AuthorName: c.AuthorName,
			Content:    c.Content,
			CreatedAt:  c.CreatedAt,
			UpdatedAt:  c.UpdatedAt,
		})
	}

	result.CommentList = commentDetailList

	return result
}
