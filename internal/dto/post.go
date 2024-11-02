package dto

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type PostReq struct {
	Title   string `json:"title" validate:"required,min=5,max=50"`
	Content string `json:"content" validate:"required,min=5,max=1000"`
}

type PostResp struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`
}

type PostDetailResp struct {
	ID       uuid.UUID        `json:"id"`
	Title    string           `json:"title"`
	Content  string           `json:"content"`
	UserID   uuid.UUID        `json:"userId"`
	Comments []models.Comment `json:"comments"`
}

func (r *PostReq) ToModel() *models.Post {
	return &models.Post{
		Title:   r.Title,
		Content: r.Content,
	}
}

func FromPost(post *models.Post) *PostResp {
	return &PostResp{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}
}

func FromPostDetail(post *models.Post) *PostDetailResp {
	return &PostDetailResp{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Comments: post.Comments,
	}
}

func FromPostList(posts []*models.Post) []*PostDetailResp {
	resp := make([]*PostDetailResp, len(posts))
	for i, post := range posts {
		resp[i] = FromPostDetail(post)
	}
	return resp
}
