package reduction_domain

import (
	money_domain "impots/internal/modules/impots/domain/money"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

type Reduction interface {
	Apply(revenu money_domain.Revenu, taxe taxe_domain.Taxe) taxe_domain.Taxe
}
