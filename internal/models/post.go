package models

import "github.com/google/uuid"

type Post struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   uuid.UUID `json:"userId"`
	Comments []Comment
}
