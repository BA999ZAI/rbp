package repository

import (
	"context"

	"rbp/internal/db/sqlc"
	"rbp/internal/models"
)

type UserRepository struct {
	querier *sqlc.Queries
}

func NewUserRepository(querier *sqlc.Queries) *UserRepository {
	return &UserRepository{querier: querier}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (int32, error) {
	res, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams{
		Email:       user.Email,
		Password:    user.Password,
		CompanyName: user.CompanyName,
		Inn:         user.INN,
		Role:        user.Role,
	})
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.querier.GetUserByEmail(ctx, sqlc.GetUserByEmailParams{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:          user.ID,
		Email:       user.Email,
		Password:    user.Password,
		CompanyName: user.CompanyName,
		INN:         user.Inn,
		Role:        user.Role,
		CreatedAt:   *user.CreatedAt,
	}, nil
}
