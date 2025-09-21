package domain

import "fmt"

// factory
type TaxReductionBasicInfo struct {
	ReductionType  string
	ReductionValue float64
}

func CreateTaxReductionCreator(trbi *TaxReductionBasicInfo) (TaxReduction, error) {
	if trbi == nil {
		trt, _ := NewReductionType("FIXE")
		return NewTaxReduction(trt, 0)
	}
	trt, err := NewReductionType(trbi.ReductionType)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return TaxReduction{}, nil
	}
	return NewTaxReduction(trt, trbi.ReductionValue)
}
