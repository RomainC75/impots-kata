package domain

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	payedTaxe Taxe
}

func NewUser(uuid uuid.UUID) *User {
	return &User{
		id: uuid,
	}
}

func (u *User) PayTaxe(taxe Taxe) {
	u.payedTaxe = u.payedTaxe.Add(taxe)
}
