package domain

import "errors"

const (
	ERROR_NEGATIVE_VALEUR = "revenu valeur cannot be negative"
)

type Revenu struct {
	valeur float64
}

func NewRevenu(valeur float64) (Revenu, error) {
	if valeur < 0 {
		return Revenu{}, errors.New(ERROR_NEGATIVE_VALEUR)
	}
	return Revenu{valeur: valeur}, nil
}

func (r Revenu) Substract(montant Montant) (Revenu, error) {
	return NewRevenu(r.valeur - montant.valeur)
}

func (r Revenu) CanSubstract(montant Montant) bool {
	return r.valeur > montant.valeur
}

func (r Revenu) ToMontant() Montant {
	return NewMontant(r.valeur)
}

func (r Revenu) Less(other Montant) bool {
	return r.valeur-other.valeur < 0
}
