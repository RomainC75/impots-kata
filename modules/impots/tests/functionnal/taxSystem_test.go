package functionnal

import (
	"fmt"
	"impots/modules/impots/applications"
	"impots/modules/impots/domain"
	"impots/modules/impots/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxSystem(t *testing.T) {
	for _, tc := range tests.Tcs {
		t.Run(fmt.Sprintf("taxSyste %f -> %f\n", tc.Revenu, tc.ExpectedTaxeToPay), func(t *testing.T) {
			taxSystem := taxSystemTestDriver(tc.AlreadyPayed)

			cisRequest := applications.CalculateImpotsServiceRequest{
				Payslip:    domain.NewRevenu(tc.Revenu),
				Reductions: tc.Reductions,
			}
			response, err := taxSystem.CalculateTax(cisRequest)
			assert.Nil(t, err)
			assert.Equal(t, tc.ExpectedTaxeToPay, response.ToBePayedTaxes)
			assert.Equal(t, tc.ExpectedTaxeBase, response.TaxeBase)
		})

	}
}
