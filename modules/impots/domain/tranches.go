package domain

type Tranches struct {
	tranches []Tranche
}

func NewTranches(revenu Montant) Tranches {
	t1, _ := NewTaxe(0)
	t2, _ := NewTaxe(10)
	t3, _ := NewTaxe(18)
	t4, _ := NewTaxe(25)
	t5, _ := NewTaxe(30)

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

func (t Tranches) CalculateTranches(revenue Montant) Montant {
	totalTaxe := NewMontant(0)
	for _, t := range t.tranches {
		totalTaxe = totalTaxe.Add(t.CalculateTrancheTaxe(revenue, &totalTaxe))
	}
	return totalTaxe
}
