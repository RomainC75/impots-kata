package reduction_domain

import (
	"errors"
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

var (
	ErrInvalidPercent          = errors.New("invalid percentage")
	PERCENT_TAX_REDUCTION_TYPE = "PERCENT"
)

type PercentTaxReduction struct {
	reductionRate  float64
	applicableFrom money_domain.Revenu
}

func NewPercentTaxReduction(reductionRate float64, applicableFrom money_domain.Revenu) (PercentTaxReduction, error) {
	if reductionRate < 0 || reductionRate > 100 {
		return PercentTaxReduction{}, ErrInvalidPercent
	}
	return PercentTaxReduction{
		reductionRate:  reductionRate,
		applicableFrom: applicableFrom,
	}, nil
}

func (fr PercentTaxReduction) Apply(revenu money_domain.Revenu, taxe taxe_domain.Taxe) taxe_domain.Taxe {
	rate := taxe.MultiplyByValue(fr.reductionRate / 100)
	return taxe.Sub(rate)
}
