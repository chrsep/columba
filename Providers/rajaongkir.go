package Providers

import (
	"columba/Consumers"
	"encoding/json"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const apiKeyEnv = "RAJA_ONGKIR_KEY"
const costEndpointUrl = "https://api.rajaongkir.com/starter/cost"

type City struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type Cost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type Service struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type Query struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
type Result struct {
	Code  string    `json:"code"`
	Name  string    `json:"name"`
	Costs []Service `json:"costs"`
}
type Data struct {
	OriginDetails      City     `json:"origin_details"`
	DestinationDetails City     `json:"destination_details"`
	Results            []Result `json:"results"`
}

type RajaOngkirCostResponse struct {
	Rajaongkir Data `json:"rajaongkir"`
}

func GetShippingRates(order Consumers.Order) (shippingRates []Consumers.ShippingRate) {
	destination := GetCity(order.Destination.City)
	origin := GetCity(order.Origin.City)

	client := http.Client{}
	payload := url.Values{}
	payload.Set("origin", origin.CityId)
	payload.Set("destination", destination.CityId)
	payload.Set("weight", strconv.Itoa(order.Weight))
	payload.Set("courier", "jne")

	request, _ := http.NewRequest("POST", costEndpointUrl, strings.NewReader(payload.Encode()))
	request.Header.Set("key", os.Getenv(apiKeyEnv))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	var parsedResponse RajaOngkirCostResponse
	_ = json.Unmarshal(body, &parsedResponse)

	for _, cost := range parsedResponse.Rajaongkir.Results[0].Costs {
		for _, service := range cost.Cost {
			shippingRates = append(shippingRates, Consumers.ShippingRate{
				Currency:    "IDR",
				ServiceCode: parsedResponse.Rajaongkir.Results[0].Code + " " + cost.Service,
				ServiceName: parsedResponse.Rajaongkir.Results[0].Name,
				Description: cost.Description,
				TotalPrice:  string(service.Value),
			})
		}
	}
	return
}

func GetCity(cityName string) (result City) {
	jsonFile, err := ioutil.ReadFile("Providers/city_data.json")
	if err != nil {
		cwd, _ := os.Getwd()
		err = raven.WrapWithExtra(err, map[string]interface{}{"cwd": cwd})
		panic(err)
	}
	var cities []City
	err = json.Unmarshal(jsonFile, &cities)
	if err != nil {
		panic(err)
	}

	for _, city := range cities {
		if strings.EqualFold(city.CityName, cityName) {
			result = city
			return
		}
	}
	return
}
