package units

import (
	"impots/modules/impots/domain"
	taxe_domain "impots/modules/impots/domain/taxe"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ReductionTC = []struct {
	Name         string
	InitialTaxe  float64
	Reductions   []domain.ReductionParameters
	ExpectedTaxe float64
}{
	{
		Name:         "no reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{},
		ExpectedTaxe: 1000,
	},
	{
		Name:         "10% reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{{"PERCENT", 10}},
		ExpectedTaxe: 900,
	},
	{
		Name:         "200 fixe reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{{"FIXE", 200}},
		ExpectedTaxe: 800,
	},
	{
		Name:         "10% + 200 fixe reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{{"PERCENT", 10}, {"FIXE", 200}},
		ExpectedTaxe: 700,
	},
	{
		Name:         "(2x percents) 10% + 20%+ 200 fixe reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{{"PERCENT", 10}, {"PERCENT", 20}, {"FIXE", 200}},
		ExpectedTaxe: 600,
	},
	{
		Name:         "(2x percents + 2x fixes) 10% + 20%+ 200 fixe reduction",
		InitialTaxe:  1000,
		Reductions:   []domain.ReductionParameters{{"PERCENT", 10}, {"PERCENT", 20}, {"FIXE", 200}, {"FIXE", 100}},
		ExpectedTaxe: 500,
	},
}

func TestReductions(t *testing.T) {
	for _, tc := range ReductionTC {
		t.Run(tc.Name, func(t *testing.T) {
			reductions, err := domain.NewReductionsHandler(tc.Reductions)
			assert.Nil(t, err)

			taxe := taxe_domain.NewTaxe(tc.InitialTaxe)
			taxe = reductions.ApplyReductions(taxe)

			assert.Equal(t, tc.ExpectedTaxe, taxe.ToFloat())
		})
	}

}
