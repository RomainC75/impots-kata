package tranche_domain

import (
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

type Tranches struct {
	tranches    []Tranche
	totalRevenu money_domain.Revenu
}

func NewTranches(revenu money_domain.Revenu, entrepreneurCAAfterAbattement money_domain.Revenu) Tranches {
	t1, _ := taxe_domain.NewTaxeRate(0)
	t2, _ := taxe_domain.NewTaxeRate(10)
	t3, _ := taxe_domain.NewTaxeRate(18)
	t4, _ := taxe_domain.NewTaxeRate(25)
	t5, _ := taxe_domain.NewTaxeRate(30)

	return Tranches{
		tranches: []Tranche{
			NewTranche(0, 10_000, t1),
			NewTranche(10_000, 10_000, t2),
			NewTranche(20_000, 10_000, t3),
			NewTranche(30_000, 20_000, t4),
			NewTranche(50_000, -1, t5),
		},
		totalRevenu: revenu.Add(entrepreneurCAAfterAbattement),
	}
}

func (ts Tranches) CalculateTaxe() (taxe_domain.Taxe, money_domain.Montant) {
	totalTaxe := taxe_domain.NewTaxe(0)
	for _, t := range ts.tranches {
		totalTaxe = totalTaxe.Add(t.CalculateTrancheTaxe(ts.totalRevenu))
	}
	taxeBase := taxe_domain.BaseMontantToBeTaxedFromRevenu(ts.totalRevenu)
	return totalTaxe, taxeBase
}
