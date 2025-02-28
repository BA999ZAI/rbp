package models

import (
	"time"
)

type User struct {
	ID          int32     `json:"id"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	CompanyName string    `json:"company_name" binding:"required"`
	INN         string    `json:"inn" binding:"required,len=12"`
	Role        string    `json:"role" binding:"required,oneof=buyer supplier"`
	CreatedAt   time.Time `json:"created_at"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
