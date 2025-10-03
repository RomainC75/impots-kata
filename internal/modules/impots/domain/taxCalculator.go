package domain

import (
	"fmt"
	"impots/internal/modules/impots/domain/entrepreneur"
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	tranche_domain "impots/internal/modules/impots/domain/tranches"
	"time"
)

type TaxCalculator struct {
	prepayedTaxe        taxe_domain.Taxe
	RevenuSalarie       money_domain.Revenu
	RevenusByEntreprise []entrepreneur.RevenuByEntreprise
	reductionsHandler   reduction_domain.ReductionsHandler
	entrepreneur        entrepreneur.Entrepreneur
}

func NewTaxCalculator(
	prepayedTaxe taxe_domain.Taxe,
	revenuSalarie money_domain.Revenu,
	reductionHandler reduction_domain.ReductionsHandler,
	revenusByEntreprise []entrepreneur.RevenuByEntreprise,
	entrepreneur entrepreneur.Entrepreneur,
) TaxCalculator {
	return TaxCalculator{
		prepayedTaxe:        prepayedTaxe,
		RevenuSalarie:       revenuSalarie,
		RevenusByEntreprise: revenusByEntreprise,
		reductionsHandler:   reductionHandler,
		entrepreneur:        entrepreneur,
	}
}

func (tc TaxCalculator) CalculateTaxeToPay(now time.Time) (taxe_domain.Taxe, money_domain.Montant, error) {
	entrepreneurCAAfterAbattement, err := tc.entrepreneur.CalculateAbattement(now, tc.RevenusByEntreprise)

	if err != nil {
		return taxe_domain.Taxe{}, money_domain.Montant{}, err
	}
	fmt.Printf("salarié %f /// entrepreneur %f\n\n", tc.RevenuSalarie.ToFloat(), entrepreneurCAAfterAbattement.ToFloat())
	tranches := tranche_domain.NewTranches(tc.RevenuSalarie, entrepreneurCAAfterAbattement)
	// brut
	taxe, taxeBase := tranches.CalculateTaxe() // - prepayedTaxe
	fmt.Println("--> total taxe (salaire + entreprises) ", taxe)
	taxe = taxe.Sub(tc.prepayedTaxe)
	// - reductions
	taxe = tc.reductionsHandler.ApplyReductions(tc.RevenuSalarie, taxe)
	// - add dividendes
	// ! =>
	return taxe, taxeBase, nil
}
