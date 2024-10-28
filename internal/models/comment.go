package models

import "github.com/google/uuid"

type Comment struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`
	PostID  uuid.UUID `json:"postId"`
}
