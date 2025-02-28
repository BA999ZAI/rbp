package models

import (
	"time"
)

type Archive struct {
	ID        int32       `json:"id"`
	UserID    int32     `json:"user_id" binding:"required"`
	ProductID int32     `json:"product_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
