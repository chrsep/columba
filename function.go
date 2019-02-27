// Package p contains an HTTP Cloud Function.
package columba

import (
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"net/http"
)

type ShippingRate struct {
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	TotalPrice  string `json:"total_price"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
}

// Response that we send back to Shopifu
type CarrierServiceResponse struct {
	Rates []ShippingRate `json:"rates"`
}

// Payload that Shopify send to our API to get shipping rates.
type IncomingQuery struct {
}

// Response that we get from raja ongkir's API.
type RajaOngkirShippingRate struct {
}

func init() {
	raven.CaptureMessageAndWait("Test 1", map[string]string{"category": "logging"})
}

// Entry point called by Cloud Function
func Columba(w http.ResponseWriter, r *http.Request) {
	raven.CaptureMessageAndWait("Test 2", map[string]string{"category": "logging"})
	response := CarrierServiceResponse{[]ShippingRate{{
		"jne",
		"jne-1",
		"20000",
		"kurir jne",
		"IDR",
	}}}
	js, _ := json.Marshal(response)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(js))
}

func GetRates(query IncomingQuery) {

}

func ExtractDetails(rates RajaOngkirShippingRate) {

}
