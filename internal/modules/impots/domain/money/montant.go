package money_domain

type Montant struct {
	value float64
}

func NewMontant(value float64) Montant {
	return Montant{
		value: value,
	}
}

func (m Montant) ToFloat() float64 {
	return m.value
}

func (m Montant) MultiplyByValue(factor float64) Montant {
	return NewMontant(m.value * factor)
}

func (m Montant) IsLess(other Montant) bool {
	return m.value < other.value
}

func (m Montant) IsMore(other Montant) bool {
	return m.value < other.value
}

func (m Montant) SubByValue(otherValue float64) Montant {
	return NewMontant(m.value - otherValue)
}

func (m Montant) Add(other Montant) Montant {
	return NewMontant(m.value + other.ToFloat())
}

func (m Montant) Sub(other Montant) Montant {
	return NewMontant(m.value - other.ToFloat())
}

// func (m Montant) ToTaxe() Taxe {
// 	return NewTaxe(m.value)
// }

func (m Montant) IsNegative() bool {
	return m.value < 0
}
