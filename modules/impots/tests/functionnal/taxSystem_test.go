package functionnal

import (
	"fmt"
	"impots/modules/impots/applications"
	"impots/modules/impots/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxSystem(t *testing.T) {
	for _, tc := range tests.Tcs {
		t.Run(fmt.Sprintf("taxSyste %f -> %f\n", tc.Revenu, tc.ExpectedTaxeToPay), func(t *testing.T) {
			taxSystem := applications.NewTaxSystem()
			cisRequest := applications.CalculateImpotsServiceRequest{
				Payslip: tc.Revenu,
			}
			response := taxSystem.CalculateTax(cisRequest)
			assert.Equal(t, tc.ExpectedTaxeToPay, response)
		})

	}
}
