package money_domain

import "errors"

var ErrNegativePercent = errors.New("rate cannnot be negative")

type RevenuRate struct {
	rate float64
}

func FromPercent(rate float64) (RevenuRate, error) {
	if rate < 0 {
		return RevenuRate{}, ErrNegativePercent
	}
	return RevenuRate{
		rate: rate,
	}, nil
}

func (rr RevenuRate) CalculateAbattement(revenu Revenu) Revenu {
	return NewRevenu(revenu.value * rr.rate / 100)
}
