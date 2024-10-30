package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Salt      string
	Posts     []Post
	Comment   []Comment
}
