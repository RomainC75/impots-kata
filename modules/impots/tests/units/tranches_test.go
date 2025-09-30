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
	{20_000, 1000},
	{30_000, 2800},
	{40_000, 5300},
}

func TestTranches(t *testing.T) {
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("-> revenu %f", tc.revenu), func(t *testing.T) {
			revenu := domain.NewMontant(tc.revenu)
			tranches := domain.NewTranches(revenu)
			taxesToPay := tranches.CalculateTranches(revenu)
			assert.Equal(t, domain.NewMontant(tc.expectedTaxeToPay), taxesToPay)
		})
	}
}
