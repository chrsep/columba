// Package p contains an HTTP Cloud Function.
package columba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Response that we get from raja ongkir's API.
type RajaOngkirShippingRate struct {
	rate
}

// Entry point called by Cloud Function
func Columba(w http.ResponseWriter, r *http.Request) {
	response := CarrierService{[]ShippingRate{{
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

// Get rates from raja ongkir
func GetRates(query IncomingQuery) (RajaOngkirShippingRate, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.rajaongkir.com/starter/cost", nil)
	if err != nil {
		return RajaOngkirShippingRate{}, err
	}
	req.Header.Set("key", os.Getenv("RAJA_ONGKIR_KEY"))
	resp, err := client.Do(req)
	var result RajaOngkirShippingRate
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// extract data from raja ongkir json and return it Shopifu's requested format
//func ExtractDetails(rates RajaOngkirShippingRate) {

//}
