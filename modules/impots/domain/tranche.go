package domain

import "fmt"

type Tranche struct {
	start  float64
	mRange float64
	taxe   Taxe
}

func NewTranche(start float64, mRange float64, taxe Taxe) Tranche {
	return Tranche{
		start:  start,
		mRange: mRange,
		taxe:   taxe,
	}
}

func (t Tranche) CalculateTrancheTaxe(revenu Montant, totalTaxeMontant *Montant) Montant {
	if revenu.ToFloat() < t.start {
		return Montant{}
	}
	tranchePartToTaxe := t.ExtractTranchePart(revenu)
	return t.taxe.CalculateTaxe(tranchePartToTaxe)
}

func (t Tranche) ExtractTranchePart(revenu Montant) Montant {
	rangeToTaxe := revenu.ToFloat() - t.start

	if t.mRange != -1 && rangeToTaxe > t.mRange {
		t.DisplayTranche(t.mRange)
		return NewMontant(t.mRange)
	}
	t.DisplayTranche(rangeToTaxe)
	return (NewMontant(rangeToTaxe))
}

func (t Tranche) DisplayTranche(trancheToTaxe float64) {
	fmt.Printf("----start : %f //  mrange : %f trancheToTaxe : %f\n", t.start, t.mRange, trancheToTaxe)
}
