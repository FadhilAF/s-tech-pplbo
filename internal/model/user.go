package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password" db:"password_hash"`
	Name         string    `json:"name" db:"name"`
	Address      string    `json:"address" db:"address"`
	Phone        string    `json:"phone" db:"phone"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required,e164"`
}

type GetUserRequest struct {
	ID uuid.UUID `json:"id" query:"id" binding:"required,uuid"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Phone   string `json:"phone" form:"phone" binding:"required,e164"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id" form:"id" binding:"required,uuid"`
}
