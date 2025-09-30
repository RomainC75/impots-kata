package domain

import "github.com/google/uuid"

type Users interface {
	GetUser(userId uuid.UUID) (User, error)
}
