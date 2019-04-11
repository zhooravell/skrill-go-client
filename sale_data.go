package skrill

type SaleData struct {
	payload
}

func NewSaleData() *SaleData {
	return &SaleData{}
}
