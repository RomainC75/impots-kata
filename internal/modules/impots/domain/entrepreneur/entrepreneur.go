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
	abattements, err := createCompanyTaxeCalculators(now, e.companies, revenuByEntrepriseDetails)
	if err != nil {
		return money_domain.Revenu{}, err
	}
	revenu := money_domain.NewRevenu(0)
	for _, abattement := range abattements {
		revenu = revenu.Add(abattement)
	}
	return revenu.Round2Decimals().ToRevenu(), nil
}

func createCompanyTaxeCalculators(now time.Time, companies []Company, revenuByEntrepriseDetails []RevenuByEntreprise) ([]money_domain.Revenu, error) {
	abattements := make([]money_domain.Revenu, 0, len(revenuByEntrepriseDetails))
	for _, revenuByEntrepriseDetail := range revenuByEntrepriseDetails {
		etc := NewEntrepreneurTaxeCalculator()
		abattement, err := etc.CalculateTaxe(now, companies, revenuByEntrepriseDetail)
		if err != nil {
			return []money_domain.Revenu{}, err
		}
		abattements = append(abattements, abattement)
	}
	return abattements, nil
}
