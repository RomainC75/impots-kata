package domain

type TaxCalculator struct{}

func NewTaxCalculator() TaxCalculator {
	return TaxCalculator{}
}

func (tc TaxCalculator) CalculateTaxeToPay(user User, paySlip Revenu) Taxe {
	tranches := NewTranches()
	taxe := tranches.CalculateTaxe(paySlip)

	return taxe.Sub(user.payedTaxe)
}
