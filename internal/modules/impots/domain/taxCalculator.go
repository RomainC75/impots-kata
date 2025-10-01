package domain

import (
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	tranche_domain "impots/internal/modules/impots/domain/tranches"
)

type TaxCalculator struct {
	prepayed          taxe_domain.Taxe
	Revenu            money_domain.Revenu
	reductionsHandler reduction_domain.ReductionsHandler
}

func NewTaxCalculator(prepayed taxe_domain.Taxe, Revenu money_domain.Revenu, reductionHandler reduction_domain.ReductionsHandler) TaxCalculator {
	return TaxCalculator{
		prepayed:          prepayed,
		Revenu:            Revenu,
		reductionsHandler: reductionHandler,
	}
}

func (tc TaxCalculator) CalculateTaxeToPay() taxe_domain.Taxe {
	tranches := tranche_domain.NewTranches(tc.Revenu)
	// brut
	taxe := tranches.CalculateTaxe()
	// - prepayed
	taxe = taxe.Sub(tc.prepayed)
	// - reductions
	taxe = tc.reductionsHandler.ApplyReductions(tc.Revenu, taxe)
	return taxe
}
