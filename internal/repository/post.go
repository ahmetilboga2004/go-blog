package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) (*models.Post, error) {
	postID := uuid.New()
	query := "INSERT INTO posts (id, title, content, user_id) VALUES (?, ?, ?, ?) RETURNING *"
	err := r.DB.QueryRow(query, postID, post.Title, post.Content, post.UserID).Scan(&post.ID, &post.Title, &post.Content, &post.UserID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetAll() ([]*models.Post, error) {
	rows, err := r.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetByID(id uuid.UUID) (*models.Post, error) {
	query := `SELECT * FROM posts WHERE id = ?`
	row := r.DB.QueryRow(query, id)
	var post models.Post
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(id uuid.UUID, post *models.Post) (*models.Post, error) {
	query := `UPDATE posts SET title = ?, content = ? WHERE id = ? RETURNING *`
	row := r.DB.QueryRow(query, post.Title, post.Content, id)
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id = ?`
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}
