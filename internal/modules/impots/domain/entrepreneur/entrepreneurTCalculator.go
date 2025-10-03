package entrepreneur

import (
	"errors"
	"impots/helpers"
	money_domain "impots/internal/modules/impots/domain/money"
	"time"
)

var (
	ErrInvalidEntrepreneurTaxe = errors.New("invalid entrepreneur taxe")
)

type EntrepreneurTaxeCalculator struct {
	commercialRevenuRate money_domain.RevenuRate
	serviceRevenuRate    money_domain.RevenuRate
}

func NewEntrepreneurTaxeCalculator() EntrepreneurTaxeCalculator {
	commercialRevenuRate, _ := money_domain.FromPercent(71)
	serviceRevenuRate, _ := money_domain.FromPercent(34)
	return EntrepreneurTaxeCalculator{
		commercialRevenuRate: commercialRevenuRate,
		serviceRevenuRate:    serviceRevenuRate,
	}
}

func (et EntrepreneurTaxeCalculator) CalculateAbattement(now time.Time, companies []Company, revenu RevenuByEntreprise) (money_domain.Revenu, error) {
	foundCompanyIndex := helpers.FindIndex(companies, func(company Company) bool {
		return company.Id == revenu.CompanyId
	})
	if foundCompanyIndex == -1 {
		return money_domain.Revenu{}, errors.New("company not found")
	}

	yearDuration := time.Hour * 24 * 365
	if now.Sub(companies[foundCompanyIndex].StartedAt) < yearDuration {
		return money_domain.NewRevenu(0), nil
	}
	switch revenu.PrestationType {
	case PrestationDeService:
		return et.serviceRevenuRate.CalculateAbattement(revenu.Revenu), nil
	case PrestationCommerciale:
		return et.commercialRevenuRate.CalculateAbattement(revenu.Revenu), nil
	default:
		return money_domain.Revenu{}, ErrInvalidEntrepreneurTaxe
	}
}
