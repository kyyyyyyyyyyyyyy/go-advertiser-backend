package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleOwner      UserRole = "owner"
	RoleAdvertiser UserRole = "advertiser"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	IsActive  bool      `json:"is_active"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
