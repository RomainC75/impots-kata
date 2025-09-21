package domain

import (
	"github.com/google/uuid"
)

type Users interface {
	GetUser(id uuid.UUID) (*User, error)
}
