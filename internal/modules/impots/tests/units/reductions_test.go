package units

import (
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ReductionTC = []struct {
	Name         string
	PaySlip      float64
	InitialTaxe  float64
	Reductions   []reduction_domain.ReductionParameters
	ExpectedTaxe float64
}{
	{
		Name:         "no reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{},
		ExpectedTaxe: 1000,
	},
	{
		Name:         "10% reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{{"PERCENT", 10, 12000}},
		ExpectedTaxe: 900,
	},
	{
		Name:         "200 fixe reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{{"FIXE", 200, 0}},
		ExpectedTaxe: 800,
	},
	{
		Name:         "10% + 200 fixe reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{{"PERCENT", 10, 0}, {"FIXE", 200, 0}},
		ExpectedTaxe: 700,
	},
	{
		Name:         "(2x percents) 10% + 20%+ 200 fixe reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{{"PERCENT", 10, 0}, {"PERCENT", 20, 0}, {"FIXE", 200, 0}},
		ExpectedTaxe: 600,
	},
	{
		Name:         "(2x percents + 2x fixes) 10% + 20%+ 200 fixe reduction",
		PaySlip:      15000,
		InitialTaxe:  1000,
		Reductions:   []reduction_domain.ReductionParameters{{"PERCENT", 10, 0}, {"PERCENT", 20, 0}, {"FIXE", 200, 12_000}, {"FIXE", 100, 0}, {"FIXE", 200, 20_000}},
		ExpectedTaxe: 500,
	},
}

func TestReductions(t *testing.T) {
	for _, tc := range ReductionTC {
		t.Run(tc.Name, func(t *testing.T) {
			reductions, err := reduction_domain.NewReductionsHandler(money_domain.NewRevenu(tc.PaySlip), tc.Reductions)
			assert.Nil(t, err)

			payslipMontant := money_domain.NewRevenu(tc.PaySlip)

			taxe := taxe_domain.NewTaxe(tc.InitialTaxe)
			taxe = reductions.ApplyReductions(payslipMontant, taxe)

			assert.Equal(t, tc.ExpectedTaxe, taxe.ToFloat())
		})
	}

}
