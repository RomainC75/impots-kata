package domain

import (
	"errors"
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

func (fr PercentTaxReduction) Apply(taxe Taxe) Taxe {
	rate := taxe.MultiplyByValue(fr.reductionRate / 100)
	return taxe.Sub(rate)
}
