package application

import "impots/domain"

type TaxSystem struct {
}

func NewTaxSystem() *TaxSystem {
	return &TaxSystem{}
}

func (ts *TaxSystem) CalculateTax(paySlip float64) (float64, error) {
	revenu, err := domain.NewRevenu(paySlip)
	if err != nil {
		return 0, err
	}
	tt := domain.NewTaxeTranches(revenu).SetTranches().Calculate()

	return tt.GetTotalTaxe().ToFloat(), nil
}
