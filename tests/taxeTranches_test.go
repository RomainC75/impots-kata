package domain_test

import (
	"fmt"
	"impots/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TaxeTranchesTestCase struct {
	revenuValue     float64
	expectedTranche [5]float64
	totalMontant    float64
}

var testCases = []TaxeTranchesTestCase{
	{5000, [5]float64{5000, 0, 0, 0, 0}, 0},
	{12000, [5]float64{10000, 2000, 0, 0, 0}, 200},
	{21000, [5]float64{10000, 10000, 1000, 0, 0}, 1180},
	{31000, [5]float64{10000, 10000, 10000, 1000, 0}, 3050},
	{54000, [5]float64{10000, 10000, 10000, 20000, 4000}, 9000},
}

func testDriver(arr [5]float64, totalMontant float64) ([5]domain.Montant, domain.Montant) {
	montants := [5]domain.Montant{}
	for i, montant := range arr {
		montants[i] = domain.NewMontant(montant)
	}
	return montants, domain.NewMontant(totalMontant)
}

func TestTaxeTranches(t *testing.T) {
	fmt.Println(testCases)
	for _, tc := range testCases {

		revenu, _ := domain.NewRevenu(tc.revenuValue)
		tranches := domain.NewTaxeTranches(revenu).SetTranches()
		tranches.Display()

		expectedTranches, totalMontant := testDriver(tc.expectedTranche, tc.totalMontant)

		assert.Equal(t, expectedTranches, tranches.GetSplitRevenus())

		tranches = tranches.Calculate()
		assert.Equal(t, totalMontant, tranches.GetTotalTaxe())
	}
}
