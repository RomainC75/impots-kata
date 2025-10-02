package entrepreneur

import (
	"errors"
	money_domain "impots/internal/modules/impots/domain/money"
	"time"

	"github.com/google/uuid"
)

type EntrepreneurActivity string

var (
	ErrCompanyNotFound = errors.New("company not found")
	ErrCompanyTooYoung = errors.New("company too young")
)

const (
	PrestationDeService   EntrepreneurActivity = "PRESTATION_SERVICE"
	PrestationCommerciale EntrepreneurActivity = "PRESTATION_COMMERCIALE"
)

type RevenuByEntreprise struct {
	CompanyId      uuid.UUID
	Revenu         money_domain.Revenu
	PrestationType EntrepreneurActivity
}

// ==============================

type CompanyTaxeCalculator func(now time.Time) (money_domain.Revenu, error)

type Company struct {
	Id        uuid.UUID
	StartedAt time.Time
}

type Entrepreneur struct {
	id        uuid.UUID
	userId    uuid.UUID
	companies []Company
}

func NewEntrepreneur(id uuid.UUID, userId uuid.UUID, companies []Company) (*Entrepreneur, error) {
	return &Entrepreneur{
		id:        id,
		userId:    userId,
		companies: companies,
	}, nil
}

func (e *Entrepreneur) CalculateAbattement(now time.Time, revenuByEntrepriseDetails []RevenuByEntreprise) (money_domain.Revenu, error) {
	companyTaxeCalculators := createCompanyTaxeCalculators(e.companies, revenuByEntrepriseDetails)
	revenu := money_domain.NewRevenu(0)
	for _, companyTaxeCalculator := range companyTaxeCalculators {
		currentCompanyTaxe, err := companyTaxeCalculator(now)
		if err != nil {
			return money_domain.Revenu{}, err
		}
		revenu = revenu.Add(currentCompanyTaxe)
	}
	return revenu.Round2Decimals().ToRevenu(), nil
}

func createCompanyTaxeCalculators(companies []Company, revenuByEntrepriseDetails []RevenuByEntreprise) []CompanyTaxeCalculator {
	calculators := make([]CompanyTaxeCalculator, 0, len(revenuByEntrepriseDetails))
	for _, revenuByEntrepriseDetail := range revenuByEntrepriseDetails {
		fn := createCompanyTaxeCalculator(companies, revenuByEntrepriseDetail)
		calculators = append(calculators, fn)
	}
	return calculators
}

func createCompanyTaxeCalculator(companies []Company, revenuByEntrepriseDetail RevenuByEntreprise) func(now time.Time) (money_domain.Revenu, error) {
	return func(now time.Time) (money_domain.Revenu, error) {
		for _, company := range companies {
			if company.Id == revenuByEntrepriseDetail.CompanyId {
				yearDuration := time.Hour * 24 * 365
				if now.Sub(company.StartedAt) > yearDuration {
					etc := NewEntrepreneurTaxeCalculator()
					return etc.CalculateTaxe(revenuByEntrepriseDetail)
				}
				return money_domain.NewRevenu(0), nil
			}
		}
		return money_domain.Revenu{}, errors.New("company not found")
	}
}
