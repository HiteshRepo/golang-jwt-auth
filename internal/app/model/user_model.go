package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	FirstName    *string   `json:"first_name" validate:"required, min=2, max=100"`
	LastName     *string   `json:"last_name" validate:"required, min=2, max=100"`
	Password     *string   `json:"password" validate:"required, min=6"`
	Email        *string   `json:"email"  validate:"email, required""`
	Phone        *string   `json:"phone" validate:"required"`
	UserType     *string   `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    *uint64   `json:"created_at"`
	UpdatedAt    *uint64   `json:"updated_at"`
}