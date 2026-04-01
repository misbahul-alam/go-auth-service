package service

import (
	"errors"

	"github.com/misbahul-alam/go-auth-service/internal/model"
	"github.com/misbahul-alam/go-auth-service/internal/repository"
	"github.com/misbahul-alam/go-auth-service/internal/utils"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(name, email, password string) error {
	exitingUser, _ := s.repo.FindByEmail(email)
	if exitingUser != nil {
		return errors.New("email already exists")
	}

	hashPassword := utils.HashPassword(password)

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashPassword,
	}
	return s.repo.Create(user)
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return user, nil

}
