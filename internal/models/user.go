package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID   		uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}