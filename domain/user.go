package domain

import "github.com/google/uuid"

type User struct {
	id               uuid.UUID
	alreadyPayedTaxe Montant
	revenu           Revenu
}

func NewUser(id uuid.UUID, revenu Revenu) *User {
	return &User{
		id:               id,
		alreadyPayedTaxe: NewMontant(0),
		revenu:           revenu,
	}
}

func (u *User) GetAlreadyPayedTaxe() Montant {
	return u.alreadyPayedTaxe
}

func (u *User) AddPayedTaxe(payed Montant) {
	u.alreadyPayedTaxe = u.alreadyPayedTaxe.Add(payed)
}

func (u *User) getRevenu() Revenu {
	return u.revenu
}
