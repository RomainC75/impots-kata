package entrepreneur

import (
	"errors"
	"fmt"
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

func createCompanyTaxeCalculators(companies []Company, revenuByEntrepriseDetails []RevenuByEntreprise) ([]CompanyTaxeCalculator, error) {
	calculators := make([]CompanyTaxeCalculator, 0, len(revenuByEntrepriseDetails))
	for _, revenuByEntrepriseDetail := range revenuByEntrepriseDetails {
		for _, company := range companies {
			if company.Id == revenuByEntrepriseDetail.CompanyId {
				fn := func(now time.Time) (money_domain.Revenu, error) {
					yearDuration := time.Hour * 24 * 365
					fmt.Println("---->", company.StartedAt, now)
					if now.Sub(company.StartedAt) > yearDuration {
						etc := NewEntrepreneurTaxeCalculator()
						return etc.CalculateTaxe(revenuByEntrepriseDetail)
					}
					return money_domain.NewRevenu(0), nil
				}
				calculators = append(calculators, fn)
				break
			}
		}
		// company not found
	}
	return calculators, nil
}

func (e *Entrepreneur) CalculateAbattement(now time.Time, revenuByEntrepriseDetails []RevenuByEntreprise) (money_domain.Revenu, error) {
	companyTaxeCalculators, err := createCompanyTaxeCalculators(e.companies, revenuByEntrepriseDetails)
	if err != nil {
		return money_domain.Revenu{}, err
	}
	revenu := money_domain.NewRevenu(0)
	for _, companyTaxeCalculator := range companyTaxeCalculators {
		currentCompanyTaxe, _ := companyTaxeCalculator(now)
		revenu = revenu.Add(currentCompanyTaxe)
		fmt.Println("-> updated entrepreneur taxe : ", revenu)
	}
	fmt.Println("resut revenu : ", revenu)
	return revenu.Round2Decimals().ToRevenu(), nil
}
