package domain

import "fmt"

type TaxeTranches struct {
	fullRevenu       Revenu
	tranches         [5]TaxeTranche
	trancheLimits    [5]float64
	revenuByTranche  [5]Montant
	totalTaxeMontant Montant
}

func NewTaxeTranches(payment *Payment) TaxeTranches {
	return TaxeTranches{
		fullRevenu:      payment.getRevenu(),
		revenuByTranche: [5]Montant{},
		tranches: [5]TaxeTranche{
			TaxeTrancheFn(NewTaxe(0)),
			TaxeTrancheFn(NewTaxe(10)),
			TaxeTrancheFn(NewTaxe(18)),
			TaxeTrancheFn(NewTaxe(25)),
			TaxeTrancheFn(NewTaxe(30)),
		},
		trancheLimits: [5]float64{
			0,
			10_000,
			20_000,
			30_000,
			50_000,
		},
	}
}

func (tt TaxeTranches) SetTranches() TaxeTranches {
	steps := [5]Montant{}
	revenuBuff := tt.fullRevenu
	for i := len(tt.trancheLimits) - 1; i >= 0; i-- {
		reste, err := revenuBuff.Substract(NewMontant(tt.trancheLimits[i]))
		if err != nil {
			steps[i] = NewMontant(0)
		} else {
			steps[i] = reste.ToMontant()
		}
		revenuBuff, _ = revenuBuff.Substract(reste.ToMontant())
	}
	tt.revenuByTranche = steps
	return tt
}

func (tt TaxeTranches) Calculate() TaxeTranches {
	total := NewMontant(0)
	for i, sr := range tt.revenuByTranche {
		total = total.Add(tt.tranches[i](sr))
	}
	tt.totalTaxeMontant = total
	return tt
}

func (tt TaxeTranches) GetRevenuByTranche() [5]Montant {
	return tt.revenuByTranche
}

func (tt TaxeTranches) GetTotalTaxe() Montant {
	return tt.totalTaxeMontant
}

func (tt TaxeTranches) Display() {
	fmt.Println("revenu", tt.fullRevenu)
	fmt.Println("----> tranches", tt.revenuByTranche)
}
