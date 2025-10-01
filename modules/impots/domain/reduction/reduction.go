package reduction_domain

import taxe_domain "impots/modules/impots/domain/taxe"

type Reduction interface {
	Apply(taxe taxe_domain.Taxe) taxe_domain.Taxe
}
