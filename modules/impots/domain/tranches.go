package domain

import money_domain "impots/modules/impots/domain/money"

type Tranches struct {
	tranches []Tranche
}

func NewTranches() Tranches {
	t1, _ := NewTaxeRate(0)
	t2, _ := NewTaxeRate(10)
	t3, _ := NewTaxeRate(18)
	t4, _ := NewTaxeRate(25)
	t5, _ := NewTaxeRate(30)

	return Tranches{
		tranches: []Tranche{
			NewTranche(0, 10_000, t1),
			NewTranche(10_000, 10_000, t2),
			NewTranche(20_000, 10_000, t3),
			NewTranche(30_000, 20_000, t4),
			NewTranche(50_000, -1, t5),
		},
	}
}

func (t Tranches) CalculateTaxe(revenue money_domain.Revenu) Taxe {
	totalTaxe := NewTaxe(0)
	for _, t := range t.tranches {
		totalTaxe = totalTaxe.Add(t.CalculateTrancheTaxe(revenue))
	}
	return totalTaxe
}
