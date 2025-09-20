package application

type TaxSystem struct {
}

func NewTaxSystem() *TaxSystem {
	return &TaxSystem{}
}

func (ts *TaxSystem) CalculateTax(paySlip float64) float64 {
	return paySlip
}
