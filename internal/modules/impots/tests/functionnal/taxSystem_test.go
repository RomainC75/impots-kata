package functionnal

import (
	"fmt"
	"impots/internal/modules/impots/applications"
	"impots/internal/modules/impots/domain/entrepreneur"
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TranchesTC struct {
	Revenu                float64
	AlreadyPayed          float64
	Reductions            []reduction_domain.ReductionParameters
	ExpectedTaxeToPay     float64
	ExpectedTaxeBase      float64
	EntrepreneurCompanies []entrepreneur.RevenuByEntreprise
}

var (
	tenPercent         = reduction_domain.ReductionParameters{"PERCENT", 10, 0}
	twentyPercent      = reduction_domain.ReductionParameters{"PERCENT", 20, 0}
	minus200           = reduction_domain.ReductionParameters{"FIXE", 200, 0}
	minus1000          = reduction_domain.ReductionParameters{"FIXE", 1000, 0}
	minus1000forRiches = reduction_domain.ReductionParameters{"FIXE", 1000, 30_000}
)

func TestTaxSystem(t *testing.T) {
	var Tcs []TranchesTC = []TranchesTC{
		{10_000, 0, []reduction_domain.ReductionParameters{}, 0, 0, []entrepreneur.RevenuByEntreprise{}},
		{15_000, 0, []reduction_domain.ReductionParameters{}, 500, 5_000, []entrepreneur.RevenuByEntreprise{}},
		{20_000, 0, []reduction_domain.ReductionParameters{}, 1000, 10_000, []entrepreneur.RevenuByEntreprise{}},
		{30_000, 0, []reduction_domain.ReductionParameters{}, 2800, 20_000, []entrepreneur.RevenuByEntreprise{}},
		{35_000, 0, []reduction_domain.ReductionParameters{}, 4050, 25_000, []entrepreneur.RevenuByEntreprise{}},
		{40_000, 0, []reduction_domain.ReductionParameters{}, 5300, 30_000, []entrepreneur.RevenuByEntreprise{}},
		{50_000, 0, []reduction_domain.ReductionParameters{}, 7800, 40_000, []entrepreneur.RevenuByEntreprise{}},
		{60_000, 0, []reduction_domain.ReductionParameters{}, 10_800, 50_000, []entrepreneur.RevenuByEntreprise{}},
		{55_000, 0, []reduction_domain.ReductionParameters{}, 9_300, 45_000, []entrepreneur.RevenuByEntreprise{}},
		// prepayed
		{10_000, 300, []reduction_domain.ReductionParameters{}, 0, 0, []entrepreneur.RevenuByEntreprise{}},
		{15_000, 200, []reduction_domain.ReductionParameters{}, 300, 5_000, []entrepreneur.RevenuByEntreprise{}},
		{20_000, 600, []reduction_domain.ReductionParameters{}, 400, 10_000, []entrepreneur.RevenuByEntreprise{}},
		{30_000, 1000, []reduction_domain.ReductionParameters{}, 1800, 20_000, []entrepreneur.RevenuByEntreprise{}},
		// prepayed + reductions
		{10_000, 300, []reduction_domain.ReductionParameters{tenPercent, minus200}, 0, 0, []entrepreneur.RevenuByEntreprise{}},
		{20_000, 600, []reduction_domain.ReductionParameters{tenPercent, minus200}, 160, 10_000, []entrepreneur.RevenuByEntreprise{}},
		{20_000, 600, []reduction_domain.ReductionParameters{tenPercent, twentyPercent, minus200}, 120, 10_000, []entrepreneur.RevenuByEntreprise{}},

		// edge cases
		{15_000, 0, []reduction_domain.ReductionParameters{minus1000, tenPercent, minus1000forRiches}, 0, 5_000, []entrepreneur.RevenuByEntreprise{}},
		{55_000, 0, []reduction_domain.ReductionParameters{minus1000, tenPercent, minus1000forRiches}, 8029, 45_000, []entrepreneur.RevenuByEntreprise{}},
		// {15_000, 0, []reduction_domain.ReductionParameters{tenPercent, twentyPercent, tenPercent, twentyPercent}, 0, 5_000},

		// with additionnal revenus from entrepreneur
		{10_000, 0, []reduction_domain.ReductionParameters{}, 0, 0, []entrepreneur.RevenuByEntreprise{
			{oldCompanyId, money_domain.NewRevenu(500), entrepreneur.PrestationCommerciale},
			{newCompanyId, money_domain.NewRevenu(500), entrepreneur.PrestationCommerciale},
		}},
	}
	for _, tc := range Tcs {
		t.Run(fmt.Sprintf("taxSyste %f -> %f\n", tc.Revenu, tc.ExpectedTaxeToPay), func(t *testing.T) {
			taxSystem := taxSystemTestDriver(tc.AlreadyPayed)

			cisRequest := applications.CalculateImpotsServiceRequest{
				RevenuSalarie: money_domain.NewRevenu(tc.Revenu),
				Reductions:    tc.Reductions,
			}
			response, err := taxSystem.CalculateTax(cisRequest)
			assert.Nil(t, err)
			assert.Equal(t, tc.ExpectedTaxeToPay, response.ToBePayedTaxes)
			assert.Equal(t, tc.ExpectedTaxeBase, response.TaxeBase)
		})

	}
}
