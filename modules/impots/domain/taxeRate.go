package domain

import "errors"

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

func (t TaxeRate) CalculateTaxe(revenu Revenu) Taxe {
	taxeMontant := revenu.MultiplyByValue(t.percent / 100)
	return taxeMontant.ToTaxe()
}
