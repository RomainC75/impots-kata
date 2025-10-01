package reduction_domain

import (
	"errors"
	taxe_domain "impots/modules/impots/domain/taxe"
)

var (
	ErrInvalidPercent          = errors.New("invalid percentage")
	PERCENT_TAX_REDUCTION_TYPE = "PERCENT"
)

type PercentTaxReduction struct {
	reductionRate float64
}

func NewPercentTaxReduction(reductionRate float64) (PercentTaxReduction, error) {
	if reductionRate < 0 || reductionRate > 100 {
		return PercentTaxReduction{}, ErrInvalidPercent
	}
	return PercentTaxReduction{
		reductionRate: reductionRate,
	}, nil
}

func (fr PercentTaxReduction) Apply(taxe taxe_domain.Taxe) taxe_domain.Taxe {
	rate := taxe.MultiplyByValue(fr.reductionRate / 100)
	return taxe.Sub(rate)
}
