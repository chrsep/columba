// Package p contains an HTTP Cloud Function.
package columba

import (
	"net/http"
)

//// Get rates from raja ongkir
//func GetRates(query IncomingQuery) (RajaOngkirShippingRate, error) {
//	client := http.Client{}
//	req, err := http.NewRequest("POST", "https://api.rajaongkir.com/starter/cost", nil)
//	if err != nil {
//		return RajaOngkirShippingRate{}, err
//	}
//	req.Header.Set("key", os.Getenv("RAJA_ONGKIR_KEY"))
//	resp, err := client.Do(req)
//	var result RajaOngkirShippingRate
//	err = json.NewDecoder(resp.Body).Decode(&result)
//	return result, nil
//}

// extract data from raja ongkir json and return it Shopifu's requested format
//func ExtractDetails(rates RajaOngkirShippingRate) {

//}
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
	//query, _ := ioutil.ReadAll(r.Body)
	//order := Consumers.ExtractOrderShopify(string(query))

}
