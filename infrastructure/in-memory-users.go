package infra

import (
	"impots/domain"

	"github.com/google/uuid"
)

type InMemoryUsers struct {
	ExpectedUser *domain.User
}

func NewInMemoryUsers() *InMemoryUsers {
	return &InMemoryUsers{}
}

func (imu *InMemoryUsers) GetUser(id uuid.UUID) (*domain.User, error) {
	return imu.ExpectedUser, nil
}
