package domain

type TaxReduction struct {
	reductionType ReductionType
	value         float64
}

func NewTaxReduction(rType ReductionType, value float64) (TaxReduction, error) {
	return TaxReduction{
		reductionType: rType,
		value:         value,
	}, nil
}

func (r TaxReduction) ApplyTo(montant Montant) Montant {
	return r.reductionType.ApplyFn()(montant, r.value)
}

func (r TaxReduction) IsType(rtype string) bool {
	return r.reductionType.rType == rtype
}

func (r TaxReduction) GetValue() float64 {
	return r.value
}
