package infra

import (
	"errors"
	"impots/internal/domain"

	"github.com/google/uuid"
)

const (
	ERROR_PAYMENT_NOT_FOUND = "payment not found"
)

type InMemoryPayments struct {
	ExpectedPayement domain.Payment
}

func NewInMemoryPayments() *InMemoryPayments {
	return &InMemoryPayments{}
}

func (imp *InMemoryPayments) ForUser(userId uuid.UUID) (domain.Payment, error) {
	if imp.ExpectedPayement.GetUserId() == userId {
		return imp.ExpectedPayement, nil
	}
	return domain.Payment{}, errors.New(ERROR_PAYMENT_NOT_FOUND)
}
