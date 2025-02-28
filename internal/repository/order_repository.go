package repository

import (
	"context"
	"rbp/internal/db/sqlc"
	"rbp/internal/models"
)

type OrderRepository struct {
	querier *sqlc.Queries
}

func NewOrderRepository(querier *sqlc.Queries) *OrderRepository {
	return &OrderRepository{querier: querier}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	return r.querier.CreateOrder(ctx, sqlc.CreateOrderParams{
		BuyerID:   order.BuyerID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
	})
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]*models.Order, error) {
	orders, err := r.querier.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	var result []*models.Order
	for _, order := range orders {
		var status string
		if err := order.Status.Scan(status); err != nil {
			return nil, err
		}

		result = append(result, &models.Order{
			ID:        order.ID,
			BuyerID:   order.BuyerID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Status:    status,
			CreatedAt: *order.CreatedAt,
		})
	}
	return result, nil
}
