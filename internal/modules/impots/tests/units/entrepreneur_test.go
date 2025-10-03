package units

import (
	"fmt"
	"impots/internal/modules/impots/domain/entrepreneur"
	money_domain "impots/internal/modules/impots/domain/money"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func EntrepreneurTestDriver() (*entrepreneur.Entrepreneur, time.Time) {
	now := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	entrepreneurId := uuid.MustParse("456e4567-e89b-12d3-a456-426614174123")
	userId := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	companies := []entrepreneur.Company{
		{
			Id:        uuid.MustParse("123e4567-e89b-12d3-a456-426614174999"),
			StartedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:        uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
			StartedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	entrepreneur, _ := entrepreneur.NewEntrepreneur(entrepreneurId, userId, companies)
	return entrepreneur, now
}

func TestEntrepreneurAbattement(t *testing.T) {
	tcs := []struct {
		name                       string
		revenueByEntrepriseDetails []entrepreneur.RevenuByEntreprise
		expectedAbattement         float64
	}{
		{
			name: "one company, prestation de service",
			revenueByEntrepriseDetails: []entrepreneur.RevenuByEntreprise{
				{
					CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
					Revenu:         money_domain.NewRevenu(10000),
					PrestationType: entrepreneur.PrestationDeService,
				},
			},
			expectedAbattement: 6600,
		},
		{
			name: "one company, prestation commerciale",
			revenueByEntrepriseDetails: []entrepreneur.RevenuByEntreprise{
				{
					CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
					Revenu:         money_domain.NewRevenu(10000),
					PrestationType: entrepreneur.PrestationCommerciale,
				},
			},
			expectedAbattement: 2900,
		},
		{
			name: "one company,prestation commerciale + prestation de service",
			revenueByEntrepriseDetails: []entrepreneur.RevenuByEntreprise{
				{
					CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
					Revenu:         money_domain.NewRevenu(10000),
					PrestationType: entrepreneur.PrestationCommerciale,
				},
				{
					CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
					Revenu:         money_domain.NewRevenu(10000),
					PrestationType: entrepreneur.PrestationDeService,
				},
			},
			expectedAbattement: 9500,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			entrepreneur, now := EntrepreneurTestDriver()
			abattement, err := entrepreneur.CalculateAbattement(now, tc.revenueByEntrepriseDetails)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedAbattement, abattement.ToFloat())
		})
	}
}

func TestErrorCompanyNotFound(t *testing.T) {
	t.Run("Company not found in RevenuByEnterprise -> error :", func(t *testing.T) {
		revenueByEntrepriseDetails := []entrepreneur.RevenuByEntreprise{
			{
				CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
				Revenu:         money_domain.NewRevenu(10000),
				PrestationType: entrepreneur.PrestationCommerciale,
			},
			{
				CompanyId:      uuid.MustParse("789e4567-e89b-12d3-a456-426614174999"),
				Revenu:         money_domain.NewRevenu(5000),
				PrestationType: entrepreneur.PrestationDeService,
			},
			{
				CompanyId:      uuid.MustParse("119e4567-e89b-12d3-a456-422614174999"),
				Revenu:         money_domain.NewRevenu(5000),
				PrestationType: entrepreneur.PrestationDeService,
			},
		}

		entrepreneur, now := EntrepreneurTestDriver()
		_, err := entrepreneur.CalculateAbattement(now, revenueByEntrepriseDetails)
		assert.Error(t, err, fmt.Errorf("company not found"))

	})
}
