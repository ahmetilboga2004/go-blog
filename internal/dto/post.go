package dto

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

// PostReq - Post isteği için kısa form
type PostReq struct {
	Title   string `json:"title" validate:"required,min=5,max=50"`
	Content string `json:"content" validate:"required,min=5,max=1000"`
}

// PostResp - Temel post yanıtı
type PostResp struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`
}

// PostDetailResp - Detaylı post yanıtı (yorumlarla)
type PostDetailResp struct {
	ID       uuid.UUID        `json:"id"`
	Title    string           `json:"title"`
	Content  string           `json:"content"`
	UserID   uuid.UUID        `json:"userId"`
	Comments []models.Comment `json:"comments"`
}

// ToModel - Post modelini oluşturur
func (r *PostReq) ToModel() *models.Post {
	return &models.Post{
		Title:   r.Title,
		Content: r.Content,
	}
}

// FromPost - Post modelinden yanıt oluşturur
func FromPost(post *models.Post) *PostResp {
	return &PostResp{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}
}

// FromPostDetail - Post modelinden detaylı yanıt oluşturur
func FromPostDetail(post *models.Post) *PostDetailResp {
	return &PostDetailResp{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		UserID:   post.UserID,
		Comments: post.Comments,
	}
}

// FromPostList - Post listesinden detaylı yanıt listesi oluşturur
func FromPostList(posts []*models.Post) []*PostDetailResp {
	resp := make([]*PostDetailResp, len(posts))
	for i, post := range posts {
		resp[i] = FromPostDetail(post)
	}
	return resp
}
