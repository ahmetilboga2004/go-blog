package dto

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type PostRequest struct {
	Title   string `json:"title" validate:"required,min=5,max=50"`
	Content string `json:"content" validate:"required,min=5,max=1000"`
}

type PostResponseWithUser struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`
}

type PostResponseWithUserAndComments struct {
	ID       uuid.UUID        `json:"id"`
	Title    string           `json:"title"`
	Content  string           `json:"content"`
	UserID   uuid.UUID        `json:"userId"`
	Comments []models.Comment `json:"comments"`
}

func (r *PostRequest) ToModel() *models.Post {
	return &models.Post{
		Title:   r.Title,
		Content: r.Content,
	}
}

func PostResponseWithUserFromModel(post *models.Post) *PostResponseWithUser {
	return &PostResponseWithUser{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}
}

func PostResponseWithUserAndCommentsFromModel(post *models.Post) *PostResponseWithUserAndComments {
	return &PostResponseWithUserAndComments{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Comments: post.Comments,
	}
}
