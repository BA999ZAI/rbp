package repository

import (
	"context"
	"rbp/internal/db/sqlc"
	"rbp/internal/models"

	"github.com/shopspring/decimal"
)

type ProductRepository struct {
	querier *sqlc.Queries
}

func NewProductRepository(querier *sqlc.Queries) *ProductRepository {
	return &ProductRepository{querier: querier}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) (int32, error) {
	res, err := r.querier.CreateProduct(ctx, sqlc.CreateProductParams{
		SupplierID:  product.SupplierID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Stock:       product.Stock,
		Photos:      product.Photos,
		Price:       decimal.NewFromFloat(product.Price),
	})

	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error) {
	products, err := r.querier.GetProductsWithFilters(ctx, sqlc.GetProductsWithFiltersParams{
		SearchQuery: "%" + req.Name + "%",
		PriceFrom:   decimal.NewFromFloat(req.PriceMin),
		PriceTo: func() decimal.Decimal {
			if req.PriceMin != 0.0 {
				return decimal.NewFromFloat(req.PriceMin)
			}
			return decimal.NewFromFloat(1000000)
		}(),
		Category: req.Category,
		InStock:  req.InStock,
	})
	if err != nil {
		return nil, err
	}

	var result []*models.Product
	for _, product := range products {
		price, _ := product.Price.Float64()
		result = append(result, &models.Product{
			ID:          product.ID,
			SupplierID:  product.SupplierID,
			Name:        product.Name,
			Description: product.Description,
			Category:    product.Category,
			Stock:       product.Stock,
			Photos:      product.Photos,
			Price:       price,
			CreatedAt:   *product.CreatedAt,
		})
	}

	return result, nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int32) (*models.Product, error) {
	product, err := r.querier.GetProductByID(ctx, sqlc.GetProductByIDParams{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	price, _ := product.Price.Float64()
	return &models.Product{
		ID:          product.ID,
		SupplierID:  product.SupplierID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       price,
		Stock:       product.Stock,
		Photos:      product.Photos,
		CreatedAt:   *product.CreatedAt,
	}, nil
}
