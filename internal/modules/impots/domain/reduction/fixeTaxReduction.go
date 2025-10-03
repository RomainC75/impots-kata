package reduction_domain

import (
	"errors"
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

var (
	ErrInvalidFixedReduction = errors.New("fixed reduction tax should be positive")
	FIXE_REDUCTION_TYPE      = "FIXE"
)

type FixeReduction struct {
	fixedTax       taxe_domain.Taxe
	applicableFrom money_domain.Revenu
}

func NewFixedReduction(fixedTax float64, applicableFrom money_domain.Revenu) (FixeReduction, error) {
	if fixedTax < 0 {
		return FixeReduction{}, ErrInvalidFixedReduction
	}
	return FixeReduction{
		fixedTax:       taxe_domain.NewTaxe(fixedTax),
		applicableFrom: applicableFrom,
	}, nil
}

func (fr FixeReduction) Apply(revenu money_domain.Revenu, taxe taxe_domain.Taxe) taxe_domain.Taxe {
	return taxe.Sub(fr.fixedTax)
}
