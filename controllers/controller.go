package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/tiagompalte/web-crawler-nf/services"
)

type Controller struct {
	service services.Service
}

func NewController(s services.Service) Controller {
	return Controller{
		service: s,
	}
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c Controller) GetBillSale(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	number := r.URL.Query().Get("number")
	uf := r.URL.Query().Get("uf")

	if url != "" {
		fmt.Println("URL", url)
		billSale, err := c.service.GetBillSaleByUrl(url)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(billSale)
	} else if number != "" && uf != "" {
		fmt.Println("Number", number)
		fmt.Println("UF", uf)
		billSale, err := c.service.GetBillSaleByNumber(number, uf)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(billSale)
	} else {
		http.Error(w, "Inform URL or Number", http.StatusInternalServerError)
	}

}
