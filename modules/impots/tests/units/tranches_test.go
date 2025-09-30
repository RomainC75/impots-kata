package units

import (
	"fmt"
	"impots/modules/impots/domain"
	"impots/modules/impots/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranches(t *testing.T) {
	for _, tc := range tests.Tcs {
		t.Run(fmt.Sprintf("-> revenu %f", tc.Revenu), func(t *testing.T) {
			revenu := domain.NewRevenu(tc.Revenu)
			tranches := domain.NewTranches()
			taxesToPay := tranches.CalculateTaxe(revenu)
			assert.Equal(t, domain.NewTaxe(tc.ExpectedTaxeToPay), taxesToPay)
		})
	}
}
