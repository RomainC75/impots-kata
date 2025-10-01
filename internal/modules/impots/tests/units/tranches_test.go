package units

import (
	"fmt"
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	tranche_domain "impots/internal/modules/impots/domain/tranches"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TranchesTC struct {
	Revenu            float64
	ExpectedTaxeToPay float64
}

var tranchesTC []TranchesTC = []TranchesTC{
	{10_000, 0},
	{15_000, 500},
	{20_000, 1000},
	{30_000, 2800},
	{35_000, 4050},
	{40_000, 5300},
	{50_000, 7800},
	{60_000, 10_800},
}

func TestTranches(t *testing.T) {
	for _, tc := range tranchesTC {
		t.Run(fmt.Sprintf("-> revenu %f", tc.Revenu), func(t *testing.T) {
			revenu := money_domain.NewRevenu(tc.Revenu)
			tranches := tranche_domain.NewTranches()
			taxesToPay := tranches.CalculateTaxe(revenu)
			assert.Equal(t, taxe_domain.NewTaxe(tc.ExpectedTaxeToPay), taxesToPay)
		})
	}
}
