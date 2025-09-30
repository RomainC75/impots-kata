package domain

type Reduction interface {
	Apply(taxe Taxe) Taxe
}
