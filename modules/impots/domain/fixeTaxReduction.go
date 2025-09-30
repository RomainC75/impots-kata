package domain

import (
	"errors"
)

var (
	ErrInvalidFixedReduction = errors.New("fixed reduction tax should be positive")
	FIXE_REDUCTION_TYPE      = "FIXE"
)

type FixeReduction struct {
	fixedTax Taxe
}

func NewFixedReduction(fixedTax float64) (FixeReduction, error) {
	if fixedTax < 0 {
		return FixeReduction{}, ErrInvalidFixedReduction
	}
	return FixeReduction{
		fixedTax: NewTaxe(fixedTax),
	}, nil
}

func (fr FixeReduction) Apply(taxe Taxe) Taxe {
	return taxe.Sub(fr.fixedTax)
}
