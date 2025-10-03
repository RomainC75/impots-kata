package user_domain

import (
	taxe_domain "impots/internal/modules/impots/domain/taxe"

	"github.com/google/uuid"
)

type User struct {
	id      uuid.UUID
	prepaid taxe_domain.Taxe
}

func NewUser(uuid uuid.UUID) *User {
	return &User{
		id: uuid,
	}
}

func (u *User) PayTaxe(taxe taxe_domain.Taxe) {
	u.prepaid = u.prepaid.Add(taxe)
}

func (u *User) GetprepaidTaxe() taxe_domain.Taxe {
	return u.prepaid
}

func (u *User) GetID() uuid.UUID {
	return u.id
}
