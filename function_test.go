package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestColumba(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: `{"message": ""}`, want: "Hello World!"},
		{body: `{"message": "Gopher"}`, want: "Hello Gopher!"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		Columba(rr, req)

		out, err := ioutil.ReadAll(rr.Result().Body)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		if got := string(out); got != test.want {
			t.Errorf("HelloHTTP(%q) = %q, want %q", test.body, got, test.want)
		}
	}
}
