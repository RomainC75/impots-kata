package infrastructure

import (
	user_domain "impots/internal/modules/impots/domain/users"

	"github.com/google/uuid"
)

type InMemoryUsers struct {
	ExpectedUser user_domain.User
}

func (imu *InMemoryUsers) GetUser(userId uuid.UUID) (user_domain.User, error) {
	return imu.ExpectedUser, nil
}

func NewInMemoryUsers() *InMemoryUsers {
	return &InMemoryUsers{}
}
