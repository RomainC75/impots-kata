package domain

import "github.com/google/uuid"

type Payment struct {
	userId           uuid.UUID
	alreadyPayedTaxe Montant
	revenu           Revenu
}

func NewPayment(userId uuid.UUID, revenu Revenu) *Payment {
	return &Payment{
		userId:           userId,
		alreadyPayedTaxe: NewMontant(0),
		revenu:           revenu,
	}
}

func (p *Payment) GetAlreadyPayedTaxe() Montant {
	return p.alreadyPayedTaxe
}

func (p *Payment) AddPayedTaxe(payed Montant) {
	p.alreadyPayedTaxe = p.alreadyPayedTaxe.Add(payed)
}

func (p *Payment) getRevenu() Revenu {
	return p.revenu
}
