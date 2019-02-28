// Package p contains an HTTP Cloud Function.
package columba

import (
	"columba/Consumers"
	"columba/Providers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	City     string
	Id       string
	Province string
}

type Order struct {
	Destination Location
	Origin      Location
	Weight      int
}

type ShippingRate struct {
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	TotalPrice  string `json:"total_price"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
}

// Entry point called by Cloud Function
func Columba(w http.ResponseWriter, r *http.Request) {
	query, _ := ioutil.ReadAll(r.Body)
	order := Consumers.ExtractOrderShopify(string(query))
	rates := Providers.GetShippingRates(order)

	response, _ := json.Marshal(rates)
	w.Header().Add("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, string(response))
}
