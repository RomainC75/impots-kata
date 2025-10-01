package functionnal

import (
	"fmt"
	"impots/internal/modules/impots/applications"
	money_domain "impots/internal/modules/impots/domain/money"
	"impots/internal/modules/impots/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxSystem(t *testing.T) {
	for _, tc := range tests.Tcs {
		t.Run(fmt.Sprintf("taxSyste %f -> %f\n", tc.Revenu, tc.ExpectedTaxeToPay), func(t *testing.T) {
			taxSystem := taxSystemTestDriver(tc.AlreadyPayed)

			cisRequest := applications.CalculateImpotsServiceRequest{
				Payslip:    money_domain.NewRevenu(tc.Revenu),
				Reductions: tc.Reductions,
			}
			response, err := taxSystem.CalculateTax(cisRequest)
			assert.Nil(t, err)
			assert.Equal(t, tc.ExpectedTaxeToPay, response.ToBePayedTaxes)
			assert.Equal(t, tc.ExpectedTaxeBase, response.TaxeBase)
		})

	}
}
