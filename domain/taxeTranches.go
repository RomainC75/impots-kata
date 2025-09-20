package domain

import "fmt"

type TaxeTranches struct {
	revenu           Revenu
	tranches         [5]TaxeTranche
	trancheLimits    [5]float64
	splitRevenu      [5]Montant
	totalTaxeMontant Montant
}

func NewTaxeTranches(revenu Revenu) TaxeTranches {
	return TaxeTranches{
		revenu:      revenu,
		splitRevenu: [5]Montant{},
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
	revenuBuff := tt.revenu
	for i := len(tt.trancheLimits) - 1; i >= 0; i-- {
		reste, err := revenuBuff.Substract(NewMontant(tt.trancheLimits[i]))
		if err != nil {
			steps[i] = NewMontant(0)
		} else {
			steps[i] = reste.ToMontant()
		}
		revenuBuff, _ = revenuBuff.Substract(reste.ToMontant())
	}
	tt.splitRevenu = steps
	return tt
}

func (tt TaxeTranches) Calculate() TaxeTranches {
	total := NewMontant(0)
	for i, sr := range tt.splitRevenu {
		total = total.Add(tt.tranches[i](sr))
	}
	tt.totalTaxeMontant = total
	return tt
}

func (tt TaxeTranches) GetSplitRevenus() [5]Montant {
	return tt.splitRevenu
}

func (tt TaxeTranches) GetTotalTaxe() Montant {
	return tt.totalTaxeMontant
}

func (tt TaxeTranches) Display() {
	fmt.Println("revenu", tt.revenu)
	fmt.Println("----> tranches", tt.splitRevenu)
}
