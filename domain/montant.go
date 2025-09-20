package domain

type Montant struct {
	valeur float64
}

func NewMontant(valeur float64) Montant {
	return Montant{
		valeur: valeur,
	}
}

func (m Montant) GetMontant() float64 {
	return m.valeur
}

func (m Montant) Add(other Montant) Montant {
	return NewMontant(m.valeur + other.valeur)
}
