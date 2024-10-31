package models

import "github.com/google/uuid"

type Post struct {
	ID       uuid.UUID
	Title    string
	Content  string
	UserID   uuid.UUID
	Comments []Comment
}
