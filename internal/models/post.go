package models

import (
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID   uuid.UUID
	Title    string
	Body     string
	User     *User
}
