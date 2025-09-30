package tests

type TranchesTC struct {
	Revenu            float64
	AlreadyPayed      float64
	ExpectedTaxeToPay float64
}

var Tcs []TranchesTC = []TranchesTC{
	{10_000, 0, 0},
	{10_000, 300, 0},
	{15_000, 0, 500},
	{15_000, 200, 300},
	{20_000, 0, 1000},
	{20_000, 600, 400},
	{30_000, 0, 2800},
	{30_000, 1000, 1800},
	{35_000, 0, 4050},
	{40_000, 0, 5300},
	{50_000, 0, 7800},
	{60_000, 0, 10_800},
}
