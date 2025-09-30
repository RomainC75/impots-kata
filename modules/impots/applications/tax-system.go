package applications

import "impots/modules/impots/domain"

type CalculateImpotsServiceRequest struct {
	Payslip float64
}

type TaxSystem struct{}

func NewTaxSystem() *TaxSystem {
	return &TaxSystem{}
}

func (cis *TaxSystem) CalculateTax(cisRequest CalculateImpotsServiceRequest) float64 {
	tranches := domain.NewTranches()
	revenu := domain.NewRevenu(cisRequest.Payslip)
	taxe := tranches.CalculateTaxe(revenu)
	return taxe.ToFloat()
}
