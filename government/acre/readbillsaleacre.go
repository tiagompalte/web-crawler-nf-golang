package acre

import (
	"net/http"
	"strings"

	"github.com/tiagompalte/web-crawler-nf/enum"
	"github.com/tiagompalte/web-crawler-nf/government"
	"github.com/tiagompalte/web-crawler-nf/util"
	"golang.org/x/net/html"
)

type ReadBillSaleAcre struct {
}

const url string = "http://www.sefaznet.ac.gov.br/nfce/qrcode?p="

func NewReadBillSaleAcre() government.ReadBillSale {
	return ReadBillSaleAcre{}
}

func (r ReadBillSaleAcre) IsGovernmentByUrl(url string) bool {
	return strings.Contains(url, "ac.gov.br")
}

func (r ReadBillSaleAcre) IsGovernmentByUf(uf string) bool {
	return strings.ToUpper(uf) == enum.Acre.String()
}

func getSession(url string) (cookie string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	cookie = res.Header.Get("Set-Cookie")
	return
}

func (r ReadBillSaleAcre) HandleUrl(url string) (billSale government.BillSale, err error) {
	cookie, err := getSession(url)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return
	}

	nodes, err := util.FindNode(
		doc,
		util.ParamFindNode{
			NodeName:   "span",
			Attributes: map[string]string{"class": "totalNumb txtMax"},
		})

	totalValue, err := util.ConvertCurrencyBrStringToFloat(util.GetContent(nodes[0]))
	if err != nil {
		return
	}

	billSale = government.BillSale{
		TotalValue: totalValue,
	}

	return
}

func (r ReadBillSaleAcre) HandleNumber(numberBillSale string) (government.BillSale, error) {
	return r.HandleUrl(url + numberBillSale)
}
