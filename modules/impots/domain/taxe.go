package domain

import "errors"

var ErrNegativePercent = errors.New("percentage cannot be negative")

type Taxe struct {
	percent float64
}

func NewTaxe(percent float64) (Taxe, error) {
	if percent < 0 {
		return Taxe{}, ErrNegativePercent
	}
	return Taxe{
		percent: percent,
	}, nil
}

func (t Taxe) CalculateTaxe(montant Montant) Montant {
	return montant.MultiplyByValue(t.percent / 100)
}
