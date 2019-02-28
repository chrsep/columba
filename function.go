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

// Entry point called by Cloud Function
func Columba(w http.ResponseWriter, r *http.Request) {
	query, _ := ioutil.ReadAll(r.Body)
	order := Consumers.ExtractOrderShopify(string(query))
	rates := Providers.GetShippingRates(order)

	response, _ := json.Marshal(rates)
	w.Header().Add("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, string(response))
}
