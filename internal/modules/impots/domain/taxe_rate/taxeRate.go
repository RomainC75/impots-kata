package taxeRate_domain

import "errors"

var ErrNegativePercent = errors.New("percentage cannot be negative")

type IRevenu interface {
	MultiplyByValue(float64) IRevenu
	ToTaxe() ITaxe
}

type ITaxe interface {
}

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

func (t TaxeRate) CalculateTaxe(revenu IRevenu) ITaxe {
	taxeMontant := revenu.MultiplyByValue(t.percent / 100)
	return taxeMontant.ToTaxe()
}
