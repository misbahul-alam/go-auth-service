package service

import (
	"github.com/misbahul-alam/go-auth-service/internal/model"
	"github.com/misbahul-alam/go-auth-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) FindByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}
