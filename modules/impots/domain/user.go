package domain

import "github.com/google/uuid"

type User struct {
	id      uuid.UUID
	prepaid Taxe
}

func NewUser(uuid uuid.UUID) *User {
	return &User{
		id: uuid,
	}
}

func (u *User) PayTaxe(taxe Taxe) {
	u.prepaid = u.prepaid.Add(taxe)
}

func (u *User) Getprepaid() Taxe {
	return u.prepaid
}
