package domain

import money_domain "impots/modules/impots/domain/money"

type TaxCalculator struct {
	prepayed          Taxe
	paySlip           money_domain.Revenu
	reductionsHandler ReductionsHandler
}

func NewTaxCalculator(prepayed Taxe, paySlip money_domain.Revenu, reductionHandler ReductionsHandler) TaxCalculator {
	return TaxCalculator{
		prepayed:          prepayed,
		paySlip:           paySlip,
		reductionsHandler: reductionHandler,
	}
}

func (tc TaxCalculator) CalculateTaxeToPay() Taxe {
	tranches := NewTranches()
	// brut
	taxe := tranches.CalculateTaxe(tc.paySlip)
	// - prepayed
	taxe = taxe.Sub(tc.prepayed)
	// - reductions
	taxe = tc.reductionsHandler.ApplyReductions(taxe)
	return taxe
}
