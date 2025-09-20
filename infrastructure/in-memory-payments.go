package infra

import (
	"impots/domain"

	"github.com/google/uuid"
)

type InMemoryPayments struct {
	ExpectedPayement domain.Payment
}

func (imp *InMemoryPayments) ForUser(userId uuid.UUID) (domain.Payment, error) {
	return imp.ExpectedPayement, nil
}
