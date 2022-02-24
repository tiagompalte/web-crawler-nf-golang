package government

type BillSale struct {
	TotalValue float64 `json:"totalValue"`
}

type ReadBillSale interface {
	IsGovernmentByUrl(url string) bool
	IsGovernmentByUf(uf string) bool
	HandleUrl(url string) (BillSale, error)
	HandleNumber(numberBillSale string) (BillSale, error)
}
