package dividendes

import (
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"

	"github.com/google/uuid"
)

const DIVIDENDE_TAXE = 30.0

type Dividende struct {
	EntrepriseId uuid.UUID
	money_domain.Revenu
}

func NewDividende(entrepriseId uuid.UUID, revenu money_domain.Revenu) Dividende {
	return Dividende{
		EntrepriseId: entrepriseId,
		Revenu:       revenu,
	}
}

func (d Dividende) GetTaxe() taxe_domain.Taxe {
	taxeRate, _ := taxe_domain.NewTaxeRate(30)
	return taxeRate.CalculateTaxe(d.Revenu)
}
