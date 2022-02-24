package services

import (
	"errors"

	"github.com/tiagompalte/web-crawler-nf/government"
)

type Service struct {
	readBillSales []government.ReadBillSale
}

func NewService(r []government.ReadBillSale) Service {
	return Service{
		readBillSales: r,
	}
}

func getImplementsByUrl(readBillSales []government.ReadBillSale, url string) (government.ReadBillSale, error) {
	for _, read := range readBillSales {
		if read.IsGovernmentByUrl(url) {
			return read, nil
		}
	}
	return nil, errors.New("ERROR: implements not found")
}

func getImplementsByUf(readBillSales []government.ReadBillSale, uf string) (government.ReadBillSale, error) {
	for _, read := range readBillSales {
		if read.IsGovernmentByUf(uf) {
			return read, nil
		}
	}
	return nil, errors.New("ERROR: implements not found")
}

func (s Service) GetBillSaleByUrl(url string) (billSale government.BillSale, err error) {
	r, err := getImplementsByUrl(s.readBillSales, url)
	if err != nil {
		return
	}
	billSale, err = r.HandleUrl(url)
	return
}

func (s Service) GetBillSaleByNumber(number string, uf string) (billSale government.BillSale, err error) {
	r, err := getImplementsByUf(s.readBillSales, uf)
	if err != nil {
		return
	}
	billSale, err = r.HandleNumber(number)
	return
}
