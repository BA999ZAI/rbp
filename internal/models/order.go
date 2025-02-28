package models

import (
	"time"
)

type Order struct {
	ID        int32     `json:"id"`
	BuyerID   int32     `json:"buyer_id" binding:"required"`
	ProductID int32     `json:"product_id" binding:"required"`
	Quantity  int32     `json:"quantity" binding:"required,gt=0"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
