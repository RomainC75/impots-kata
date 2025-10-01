package taxe_domain

import (
	"errors"
	money_domain "impots/internal/modules/impots/domain/money"
)

var ErrNegativePercent = errors.New("percentage cannot be negative")

type TaxeRate struct {
	percent float64
}

func NewTaxeRate(percent float64) (TaxeRate, error) {
	if percent < 0 {
		return TaxeRate{}, ErrNegativePercent
	}
	return TaxeRate{
		percent: percent,
	}, nil
}

func (t TaxeRate) CalculateTaxe(revenu money_domain.Revenu) Taxe {
	taxeMontant := revenu.MultiplyByValue(t.percent / 100)
	return NewTaxeFromMontant(taxeMontant)
}
