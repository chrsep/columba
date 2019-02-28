// Package p contains an HTTP Cloud Function.
package columba

import (
	"columba/Consumers"
	"columba/Providers"
	"encoding/json"
	"fmt"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

// Entry point called by Cloud Function
func Columba(w http.ResponseWriter, r *http.Request) {
	var response []byte
	err, _ := raven.CapturePanic(func() {
		query, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		order := Consumers.ExtractOrderShopify(string(query))
		rates := Providers.GetShippingRates(order)
		response, _ = json.Marshal(rates)

		w.Header().Add("Content-Type", "application/json")
	}, nil)

	if err != nil {
		errorResponse := ErrorResponse{
			Status:      "500",
			Description: fmt.Sprintf("Error: %v", err),
		}
		response, _ = json.Marshal(errorResponse)
		_, _ = fmt.Fprint(w, string(response))
	} else {
		_, _ = fmt.Fprint(w, string(response))
	}
}
