package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) Create(comment *models.Comment) (*models.Comment, error) {
	commentID := uuid.New()
	query := `INSERT INTO comments (id, content, user_id, post_id) VALUES (?, ?, ?, ?) RETURNING *`
	row := r.DB.QueryRow(query, commentID, comment.Content, comment.UserID, comment.PostID)
	if err := row.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID); err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) GetAll() ([]*models.Comment, error) {
	rows, err := r.DB.Query("SELECT * FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) GetByID(id uuid.UUID) (*models.Comment, error) {
	query := `SELECT * FROM comments WHERE id = ?`
	row := r.DB.QueryRow(query, id)
	var comment models.Comment
	if err := row.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) Update(id uuid.UUID, comment *models.Comment) (*models.Comment, error) {
	query := "UPDATE comments SET content = ? WHERE id = ? RETURNONG *"
	row := r.DB.QueryRow(query, comment.Content, id)
	if err := row.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) Delete(id uuid.UUID) error {
	result, err := r.DB.Exec("DELETE FROM comments WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("comment not found")
	}
	return nil
}
