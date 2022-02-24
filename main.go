package main

import (
	"fmt"
	"net/http"

	"github.com/tiagompalte/web-crawler-nf/controllers"
	"github.com/tiagompalte/web-crawler-nf/government"
	"github.com/tiagompalte/web-crawler-nf/government/acre"
	"github.com/tiagompalte/web-crawler-nf/services"
)

func main() {
	r := []government.ReadBillSale{acre.NewReadBillSaleAcre()}
	service := services.NewService(r)
	controller := controllers.NewController(service)

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/bill-sale", controller.GetBillSale)

	fmt.Println("Listening: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
