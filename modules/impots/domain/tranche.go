package domain

import (
	"fmt"
	money_domain "impots/modules/impots/domain/money"
)

type Tranche struct {
	start  float64
	mRange float64
	taxe   TaxeRate
}

func NewTranche(start float64, mRange float64, taxe TaxeRate) Tranche {
	return Tranche{
		start:  start,
		mRange: mRange,
		taxe:   taxe,
	}
}

func (t Tranche) CalculateTrancheTaxe(revenu money_domain.Revenu) Taxe {
	if revenu.ToFloat() < t.start {
		return Taxe{}
	}
	trancheTaxe := t.ExtractTranchePart(revenu)
	return t.taxe.CalculateTaxe(money_domain.Revenu(trancheTaxe))
}

func (t Tranche) ExtractTranchePart(revenu money_domain.Revenu) Taxe {
	rangeToTaxe := revenu.ToFloat() - t.start

	if t.mRange != -1 && rangeToTaxe > t.mRange {
		t.DisplayTranche(t.mRange)
		return NewTaxe(t.mRange)
	}
	t.DisplayTranche(rangeToTaxe)
	return (NewTaxe(rangeToTaxe))
}

func (t Tranche) DisplayTranche(trancheToTaxe float64) {
	fmt.Printf("----start : %f //  mrange : %f trancheToTaxe : %f\n", t.start, t.mRange, trancheToTaxe)
}
