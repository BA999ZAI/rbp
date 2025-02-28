package service

import (
	"context"
	"rbp/internal/models"
	"rbp/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) error {
	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderService) GetOrders(ctx context.Context) ([]*models.Order, error) {
	return s.repo.GetOrders(ctx)
}
