package domain

import (
	"errors"
	taxe_domain "impots/modules/impots/domain/taxe"
)

var (
	ErrInvalidFixedReduction = errors.New("fixed reduction tax should be positive")
	FIXE_REDUCTION_TYPE      = "FIXE"
)

type FixeReduction struct {
	fixedTax taxe_domain.Taxe
}

func NewFixedReduction(fixedTax float64) (FixeReduction, error) {
	if fixedTax < 0 {
		return FixeReduction{}, ErrInvalidFixedReduction
	}
	return FixeReduction{
		fixedTax: taxe_domain.NewTaxe(fixedTax),
	}, nil
}

func (fr FixeReduction) Apply(taxe taxe_domain.Taxe) taxe_domain.Taxe {
	return taxe.Sub(fr.fixedTax)
}
