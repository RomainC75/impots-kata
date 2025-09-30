package infrastructure

import (
	"impots/modules/impots/domain"

	"github.com/google/uuid"
)

type InMemoryUsers struct {
	ExpectedUser domain.User
}

func (imu *InMemoryUsers) GetUser(userId uuid.UUID) (domain.User, error) {
	return imu.ExpectedUser, nil
}

func NewInMemoryUsers() *InMemoryUsers {
	return &InMemoryUsers{}
}
