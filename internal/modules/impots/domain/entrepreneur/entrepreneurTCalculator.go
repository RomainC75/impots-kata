package entrepreneur

import (
	"errors"
	money_domain "impots/internal/modules/impots/domain/money"
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

func (et EntrepreneurTaxeCalculator) CalculateTaxe(revenu RevenuByEntreprise) (money_domain.Revenu, error) {
	switch revenu.PrestationType {
	case PrestationDeService:
		return et.serviceRevenuRate.CalculateAbattement(revenu.Revenu), nil
	case PrestationCommerciale:
		return et.commercialRevenuRate.CalculateAbattement(revenu.Revenu), nil
	default:
		return money_domain.Revenu{}, ErrInvalidEntrepreneurTaxe
	}
}
