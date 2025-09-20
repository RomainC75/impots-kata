package domain_test

import (
	"fmt"
	"impots/domain"
	"testing"

	"github.com/google/uuid"
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
		expectedTranches, totalMontant := testDriver(tc.expectedTranche, tc.totalMontant)

		revenu, _ := domain.NewRevenu(tc.revenuValue)
		userUuid := uuid.MustParse("45c971a4-5aeb-40e8-ba51-0f6698e92528")
		user := domain.NewUser(userUuid, revenu)

		tranches := domain.NewTaxeTranches(user).SetTranches()
		tranches.Display()

		assert.Equal(t, expectedTranches, tranches.GetRevenuByTranche())

		tranches = tranches.Calculate()
		assert.Equal(t, totalMontant, tranches.GetTotalTaxe())
	}
}
