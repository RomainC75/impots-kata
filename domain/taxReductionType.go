package domain

import (
	"errors"
)

const (
	ERROR_INVALID_REDUCTION_TYPE = "invalid reduction type"
)

type ReductionType struct {
	rType string
	algo  func(montant Montant, reductionValue float64) Montant
}

func NewReductionType(str string) (ReductionType, error) {
	if str != "FIXE" && str != "PERCENT" {
		return ReductionType{}, errors.New(ERROR_INVALID_REDUCTION_TYPE)
	}
	rt := ReductionType{rType: str}
	return rt.setAlgo(), nil
}

func (rt ReductionType) setAlgo() ReductionType {
	if rt.rType == "FIXE" {
		rt.algo = rt.ApplyFixeReduction()
		return rt
	}
	rt.algo = rt.ApplyPercentReduction()
	return rt
}

func (rt ReductionType) ApplyFn() func(montant Montant, reductionValue float64) Montant {
	return rt.algo
}

func (rt ReductionType) ApplyFixeReduction() func(montant Montant, reductionValue float64) Montant {
	return func(montant Montant, reductionValue float64) Montant {
		res := montant.ToFloat() - reductionValue
		if res < 0 {
			return NewMontant(0)
		}
		return NewMontant(res)
	}
}

func (rt ReductionType) ApplyPercentReduction() func(montant Montant, reductionValue float64) Montant {
	return func(montant Montant, reductionValue float64) Montant {
		return montant.Sub(NewMontant(montant.ToFloat() * reductionValue / 100))
	}
}
