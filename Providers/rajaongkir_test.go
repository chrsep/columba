package Providers

import (
	"columba"
	"testing"
)

func TestGetCity(t *testing.T) {
	tests := []struct {
		cityName string
		want     string
	}{
		{cityName: "Aceh Barat", want: "1"},
		{cityName: "Aceh Besar", want: "3"},
		{cityName: "Jakarta Barat", want: "151"},
		{cityName: "jakarta selatan", want: "153"},
	}

	for _, test := range tests {
		city := GetCity(test.cityName)
		if city.CityId != test.want {
			t.Errorf("TestGetCity expects %v, get %v", test.want, city.CityId)
		}
	}
}

func TestGetShippingRates(t *testing.T) {
	order := columba.Order{
		Weight: 2000,
		Origin: columba.Location{
			City:     "Jakarta Barat",
			Id:       "",
			Province: "Jabodetabek",
		},
		Destination: columba.Location{
			City:     "Tangerang Selatan",
			Province: "Banten",
			Id:       "",
		},
	}

	rates := GetShippingRates(order)
	if len(rates) == 0 {
		t.Errorf("Rates should not returned empty")
	}
}