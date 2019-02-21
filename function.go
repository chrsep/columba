// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CarrierResponse struct {
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	TotalPrice  string `json:"total_price"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
}

type Respone struct {
	Rates []CarrierResponse `json:"rates"`
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Calculate(w http.ResponseWriter, r *http.Request) {
	response := Respone{[]CarrierResponse{{
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
