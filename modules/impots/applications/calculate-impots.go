package applications

type CalculateImpotsServiceRequest struct {
	Payslip float64
}

type CalculateImpotsService struct{}

func NewCalculateImpotsService() *CalculateImpotsService {
	return &CalculateImpotsService{}
}

func (cis *CalculateImpotsService) Calculate(cisRequest CalculateImpotsServiceRequest) {

}
