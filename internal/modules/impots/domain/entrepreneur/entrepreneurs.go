package entrepreneur

import "github.com/google/uuid"

type Entrepreneurs interface {
	GetEntrepreneur(id uuid.UUID) (Entrepreneur, error)
}
