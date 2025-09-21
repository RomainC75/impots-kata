package domain

type TaxeTranche = func(montant Montant) Montant

func TaxeTrancheFn(taxe Taxe) func(montant Montant) Montant {
	return func(montant Montant) Montant {
		return taxe.ApplyToMontant(montant)
	}
}
