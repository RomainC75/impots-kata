package domain

type TaxCalculator struct{}

func NewTaxCalculator() TaxCalculator {
	return TaxCalculator{}
}

func (tc TaxCalculator) CalculateTaxeToPay(user User, paySlip float64) Taxe {
	tranches := NewTranches()
	revenu := NewRevenu(paySlip)
	taxe := tranches.CalculateTaxe(revenu)

	return taxe.Sub(user.payedTaxe)
}
