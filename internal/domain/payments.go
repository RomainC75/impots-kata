package domain

import "github.com/google/uuid"

type Payments interface {
	ForUser(userId uuid.UUID) (Payment, error)
}
