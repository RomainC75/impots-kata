package units

import (
	"fmt"
	"impots/modules/impots/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TranchesTC struct {
	revenu            float64
	expectedTaxeToPay float64
}

var tcs []TranchesTC = []TranchesTC{
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
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("-> revenu %f", tc.revenu), func(t *testing.T) {
			revenu := domain.NewRevenu(tc.revenu)
			tranches := domain.NewTranches()
			taxesToPay := tranches.CalculateTaxe(revenu)
			assert.Equal(t, domain.NewTaxe(tc.expectedTaxeToPay), taxesToPay)
		})
	}
}
