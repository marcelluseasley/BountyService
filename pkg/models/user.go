package models

import (
	"time"
)

type User struct {
	ID           int    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	JWT          string    `json:"jwt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserService interface {
	CreateUser() error 
	DeleteUser() error 
	Login() (string, error)
}
