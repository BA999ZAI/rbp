package service

import (
	"errors"
	"rbp/internal/repository"
	"rbp/internal/utils"
	"rbp/pkg/cache"
	"time"
)

type AuthService struct {
	repo        *repository.UserRepository
	redisClient *cache.RedisClient
}

func NewAuthService(repo *repository.UserRepository, redisClient *cache.RedisClient) *AuthService {
	return &AuthService{repo: repo, redisClient: redisClient}
}

func (s *AuthService) ValidateToken(token string) (int, error) {
	claims, err := utils.ParseJWT(token)
	if err != nil {
		return 0, err
	}
	userID := claims["user_id"].(float64)
	return int(userID), nil
}

func (s *AuthService) GenerateRefreshToken(userID int) (string, error) {
	token := utils.GenerateRefreshToken()
	err := s.redisClient.Set(token, userID, 7*24*time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	userID, err := s.redisClient.Get(refreshToken)
	if err != nil {
		return "", err
	}
	if userID == 0 {
		return "", errors.New("invalid refresh token")
	}

	return utils.GenerateJWT(int32(userID))
}
