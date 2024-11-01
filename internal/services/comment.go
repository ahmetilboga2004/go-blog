package services

import (
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type commentService struct {
	commentRepo interfaces.CommentRepository
}

func NewcommentService(commentRepo interfaces.CommentRepository) interfaces.CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (s *commentService) CreateComment(userId uuid.UUID, comment *models.Comment) (*models.Comment, error) {
	comment.UserID = userId
	comment, err := s.commentRepo.Create(comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) GetCommentByID(id uuid.UUID) (*models.Comment, error) {
	comment, err := s.commentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) GetAllComments() ([]*models.Comment, error) {
	comments, err := s.commentRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *commentService) UpdateComment(userId, commentId uuid.UUID, comment *models.Comment) (*models.Comment, error) {
	commentCheck, err := s.commentRepo.GetByID(commentId)
	if err != nil {
		return nil, errors.New("comment not found")
	}
	if commentCheck.UserID != userId {
		return nil, errors.New("unauthorized")
	}

	comment, err = s.commentRepo.Update(commentId, comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) DeleteComment(userId, commentId uuid.UUID) error {
	checkComment, err := s.commentRepo.GetByID(commentId)
	if err != nil {
		return err
	}
	if checkComment.UserID != userId {
		return errors.New("unauthorized")
	}
	if err := s.commentRepo.Delete(commentId); err != nil {
		return err
	}
	return nil
}
