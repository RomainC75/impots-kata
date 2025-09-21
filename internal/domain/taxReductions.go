package domain

import (
	"fmt"
)

type TaxReductionBasicInfo struct {
	ReductionType  string
	ReductionValue float64
}

type TaxReductions struct {
	reductions []TaxReduction
}

// PERCENT X 1
// PERCENT AVANT FIXE

func NewTaxeReductions(trbis []TaxReductionBasicInfo) (TaxReductions, error) {
	reductions := []TaxReduction{}
	for _, trbi := range trbis {
		trt, err := NewReductionType(trbi.ReductionType)
		if err != nil {
			fmt.Println("error : ", err.Error())
			return TaxReductions{}, nil
		}
		taxeReduction, _ := NewTaxReduction(trt, trbi.ReductionValue)
		reductions = append(reductions, taxeReduction)

	}
	return TaxReductions{
		reductions: reductions,
	}, nil
}

func (trs TaxReductions) ApplyTo(montant Montant) Montant {
	biggestPercentIndex := trs.findBiggestPercentIndex()
	if biggestPercentIndex != -1 {
		montant = trs.reductions[biggestPercentIndex].ApplyTo(montant)
	}
	for _, reduction := range trs.reductions {
		if reduction.IsType("FIXE") {
			montant = reduction.ApplyTo(montant)
		}
	}
	if montant.ToFloat() < 0 {
		return NewMontant(0)
	}
	return montant
}

func (trs TaxReductions) findBiggestPercentIndex() int {
	var biggestP float64 = -1
	index := -1

	for i, reduction := range trs.reductions {
		if reduction.IsType("PERCENT") && reduction.GetValue() > biggestP {
			biggestP = reduction.GetValue()
			index = i
		}
	}
	return index
}
