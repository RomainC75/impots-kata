package money_domain

type Revenu struct {
	Montant
}

func NewRevenu(value float64) Revenu {
	return Revenu{
		Montant: NewMontant(value),
	}
}

func (r Revenu) Add(other Revenu) Revenu {
	totalMontant := r.Montant.Add(other.Montant)
	return NewRevenu(totalMontant.value)
}
