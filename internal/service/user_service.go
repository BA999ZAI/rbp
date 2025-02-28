package service

import (
	"context"
	"errors"

	"rbp/internal/models"
	"rbp/internal/repository"
	"rbp/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *models.User) (int32, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid password")
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshRefreshToken, err := utils.GenerateRefresJWT(user.ID)
	if err != nil {
		return "", "", err
	}

	return token, refreshRefreshToken, nil
}
