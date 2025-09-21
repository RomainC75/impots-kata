package domain

import "github.com/google/uuid"

type User struct {
	id uuid.UUID
}

func NewUser(id uuid.UUID, revenu Revenu) *User {
	return &User{
		id: id,
	}
}
