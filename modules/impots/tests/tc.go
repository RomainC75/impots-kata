package tests

type TranchesTC struct {
	Revenu            float64
	AlreadyPayed      float64
	ExpectedTaxeToPay float64
	ExpectedTaxeBase  float64
}

var Tcs []TranchesTC = []TranchesTC{
	{10_000, 0, 0, 0},
	{10_000, 300, 0, 0},
	{15_000, 0, 500, 5_000},
	{15_000, 200, 300, 5_000},
	{20_000, 0, 1000, 10_000},
	{20_000, 600, 400, 10_000},
	{30_000, 0, 2800, 20_000},
	{30_000, 1000, 1800, 20_000},
	{35_000, 0, 4050, 25_000},
	{40_000, 0, 5300, 30_000},
	{50_000, 0, 7800, 40_000},
	{60_000, 0, 10_800, 50_000},
}
