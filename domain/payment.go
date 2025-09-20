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

func (p *Payment) GetTaxableBase() Montant {
	if p.revenu.Less(NewMontant(10_000)) {
		return NewMontant(0)
	}
	return NewMontant(p.revenu.valeur - 10_000)
}

func (p *Payment) AddPayedTaxe(payed Montant) {
	p.alreadyPayedTaxe = p.alreadyPayedTaxe.Add(payed)
}

func (p *Payment) getRevenu() Revenu {
	return p.revenu
}

func (p *Payment) GetUserId() uuid.UUID {
	return p.userId
}
