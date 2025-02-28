package service

import (
	"context"
	"rbp/internal/models"
	"rbp/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (int32, error) {
	return s.repo.CreateProduct(ctx, product)
}

func (s *ProductService) GetProducts(ctx context.Context, req *models.ProductFilters) ([]*models.Product, error) {
	return s.repo.GetProducts(ctx, req)
}

func (s *ProductService) GetProductByID(ctx context.Context, id int32) (*models.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}
