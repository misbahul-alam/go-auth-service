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

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}
	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid password")
	}
	AccessToken := utils.GenerateAccessToken(user.ID, string(user.Role))
	RefreshToken := utils.GenerateRefreshToken(user.ID, string(user.Role))
	return AccessToken, RefreshToken, nil

}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}
	if claims == nil {
		return "", errors.New("invalid refresh token")
	}
	if claims.Type != "refresh" {
		return "", errors.New("invalid refresh token")
	}
	AccessToken := utils.GenerateAccessToken(claims.UserID, claims.Role)
	return AccessToken, nil
}
