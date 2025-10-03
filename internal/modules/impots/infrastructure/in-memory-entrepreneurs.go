package infrastructure

import (
	"impots/internal/modules/impots/domain/entrepreneur"

	"github.com/google/uuid"
)

type InMemoryEntrepreneurs struct {
	ExpectedEntrepreneur entrepreneur.Entrepreneur
}

func NewInMemoryEntrepreneurs() *InMemoryEntrepreneurs {
	return &InMemoryEntrepreneurs{}
}

func (ine *InMemoryEntrepreneurs) GetEntrepreneur(id uuid.UUID) (entrepreneur.Entrepreneur, error) {
	return ine.ExpectedEntrepreneur, nil
}
