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
	reduction       *domain.TaxReductionBasicInfo
}

var testCases = []TaxeTranchesTestCase{
	// NO REDUCTION
	{revenuValue: 5000, expectedTranche: [5]float64{5000, 0, 0, 0, 0}, totalMontant: 0},
	{revenuValue: 12000, expectedTranche: [5]float64{10000, 2000, 0, 0, 0}, totalMontant: 200},
	{revenuValue: 21000, expectedTranche: [5]float64{10000, 10000, 1000, 0, 0}, totalMontant: 1180},
	{revenuValue: 31000, expectedTranche: [5]float64{10000, 10000, 10000, 1000, 0}, totalMontant: 3050},
	{revenuValue: 54000, expectedTranche: [5]float64{10000, 10000, 10000, 20000, 4000}, totalMontant: 9000, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 0}},

	// FIXE
	{revenuValue: 5000, expectedTranche: [5]float64{5000, 0, 0, 0, 0}, totalMontant: 0},
	{revenuValue: 12000, expectedTranche: [5]float64{10000, 2000, 0, 0, 0}, totalMontant: 100, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 100}},
	{revenuValue: 21000, expectedTranche: [5]float64{10000, 10000, 1000, 0, 0}, totalMontant: 680, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 500}},
	{revenuValue: 31000, expectedTranche: [5]float64{10000, 10000, 10000, 1000, 0}, totalMontant: 2050, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 1000}},
	{revenuValue: 54000, expectedTranche: [5]float64{10000, 10000, 10000, 20000, 4000}, totalMontant: 7000, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 2000}},

	// PERCENT
	{revenuValue: 54000, expectedTranche: [5]float64{10000, 10000, 10000, 20000, 4000}, totalMontant: 4500, reduction: &domain.TaxReductionBasicInfo{ReductionType: "PERCENT", ReductionValue: 50}},
	{revenuValue: 21000, expectedTranche: [5]float64{10000, 10000, 1000, 0, 0}, totalMontant: 1062, reduction: &domain.TaxReductionBasicInfo{ReductionType: "PERCENT", ReductionValue: 10}},
	{revenuValue: 31000, expectedTranche: [5]float64{10000, 10000, 10000, 1000, 0}, totalMontant: 2440, reduction: &domain.TaxReductionBasicInfo{ReductionType: "PERCENT", ReductionValue: 20}},

	// EDGE
	{revenuValue: 12000, expectedTranche: [5]float64{10000, 2000, 0, 0, 0}, totalMontant: 0, reduction: &domain.TaxReductionBasicInfo{ReductionType: "FIXE", ReductionValue: 300}},
	{revenuValue: 54000, expectedTranche: [5]float64{10000, 10000, 10000, 20000, 4000}, totalMontant: 0, reduction: &domain.TaxReductionBasicInfo{ReductionType: "PERCENT", ReductionValue: 120}},
}

func taxeTranchesTestDriver(tc TaxeTranchesTestCase) ([5]domain.Montant, domain.Montant, domain.TaxReduction) {

	montants := [5]domain.Montant{}
	for i, montant := range tc.expectedTranche {
		montants[i] = domain.NewMontant(montant)
	}

	taxReduction, _ := domain.CreateTaxReductionCreator(tc.reduction)

	return montants, domain.NewMontant(tc.totalMontant), taxReduction
}

func TestTaxeTranches(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("-> %d : shoud calculate tax tranches %f\n", i, tc.revenuValue), func(t *testing.T) {
			expectedTranches, totalMontant, taxReduction := taxeTranchesTestDriver(tc)

			userUuid := uuid.MustParse("45c971a4-5aeb-40e8-ba51-0f6698e92528")
			revenu, _ := domain.NewRevenu(tc.revenuValue)
			payment := domain.NewPayment(userUuid, revenu)

			tranches := domain.NewTaxeTranches(payment).SetTranches()
			// tranches.Display()

			assert.Equal(t, expectedTranches, tranches.GetRevenuByTranche())

			tranches = tranches.Calculate()
			assert.Equal(t, totalMontant, tranches.GetTotalTaxe(taxReduction))
		})
	}
}
