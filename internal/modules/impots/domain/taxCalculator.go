package domain

import (
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	tranche_domain "impots/internal/modules/impots/domain/tranches"
)

type TaxCalculator struct {
	prepayed          taxe_domain.Taxe
	paySlip           money_domain.Revenu
	reductionsHandler reduction_domain.ReductionsHandler
}

func NewTaxCalculator(prepayed taxe_domain.Taxe, paySlip money_domain.Revenu, reductionHandler reduction_domain.ReductionsHandler) TaxCalculator {
	return TaxCalculator{
		prepayed:          prepayed,
		paySlip:           paySlip,
		reductionsHandler: reductionHandler,
	}
}

func (tc TaxCalculator) CalculateTaxeToPay() taxe_domain.Taxe {
	tranches := tranche_domain.NewTranches(tc.paySlip)
	// brut
	taxe := tranches.CalculateTaxe()
	// - prepayed
	taxe = taxe.Sub(tc.prepayed)
	// - reductions
	taxe = tc.reductionsHandler.ApplyReductions(tc.paySlip, taxe)
	return taxe
}
