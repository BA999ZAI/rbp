package models

import (
	"time"
)

type Product struct {
	ID          int32     `json:"id"`
	SupplierID  int32     `json:"supplier_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description *string   `json:"description,omitempty"`
	Category    string    `json:"category"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	Stock       int32     `json:"stock" binding:"required,gte=0"`
	Photos      []string  `json:"photos"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductFilters struct {
	Name     string  `json:"name,omitempty"`
	InStock  bool    `json:"in_stock,omitempty"`
	Category string  `json:"category,omitempty"`
	PriceMin float64 `json:"price_min,omitempty"`
	PriceMax float64 `json:"price_max,omitempty"`
}
