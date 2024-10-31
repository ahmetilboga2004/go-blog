package services

import (
	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type postService struct {
	postRepo interfaces.PostRepository
}

func NewPostService(postRepo interfaces.PostRepository) interfaces.PostService {
	return &postService{
		postRepo: postRepo,
	}
}

func (s *postService) CreatePost(userId uuid.UUID, post *models.Post) (*models.Post, error) {
	post.UserID = userId
	post, err := s.postRepo.Create(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetPostByID(id uuid.UUID) (*models.Post, error) {
	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetAllPosts() ([]*models.Post, error) {
	posts, err := s.postRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *postService) UpdatePost(id uuid.UUID, post *models.Post) (*models.Post, error) {
	post, err := s.postRepo.Update(id, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) DeletePost(id uuid.UUID) error {
	if err := s.postRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
