package Consumers

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"testing"
)

func TestCalculateWeight(t *testing.T) {
	tests := []struct {
		body string
		want int
	}{
		{
			body: `{ "items": [ { "quantity": 1, "grams": 1000, } ] }`,
			want: 1000,
		}, {
			body: `{ "items": [ { "quantity": 2, "grams": 1000, } ] }`,
			want: 2000,
		}, {
			body: `{ "items": [ { "quantity": 0, "grams": 34327, } ] }`,
			want: 0,
		}, {
			body: `{ "items": [ { "quantity": 100, "grams": 54380, } ] }`,
			want: 5438000,
		},
	}

	for _, test := range tests {
		parsedItems := gjson.Get(test.body, "items")
		weight := CalculateWeight(parsedItems)
		if weight != test.want {
			t.Errorf("CalculateWeight returns %d, want %d", weight, test.want)
		}
	}
}

func TestExtractOrderShopify(t *testing.T) {
	json, err := ioutil.ReadFile("request_mock.json")
	if err != nil {
		return
	}
	tests := []struct {
		body string
		want Order
	}{
		{
			body: string(json),
			want: Order{
				Destination: Location{
					Province: "ON",
					Id:       "",
					City:     "Ottawa",
				},
				Origin: Location{
					Province: "ON",
					City:     "Ottawa",
				},
				Weight: 1000,
			},
		},
	}

	for _, test := range tests {
		order := ExtractOrderShopify(test.body)
		if order.Origin != test.want.Origin {
			t.Errorf("ExtractOrderShopify origin returns %v, want %v", order.Origin, test.want.Origin)
		}

		if order.Destination != test.want.Destination {
			t.Errorf("ExtractOrderShopify origin returns %v, want %v", order.Destination, test.want.Destination)
		}

		if order.Weight != test.want.Weight {
			t.Errorf("ExtractOrderShopify origin returns %v, want %v", order.Weight, test.want.Weight)
		}
	}
}
