package Consumers

import (
	"github.com/tidwall/gjson"
)

type ShopifyCarrierServiceResponse struct {
	Rates []ShippingRate `json:"rates"`
}
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

func CalculateWeight(items gjson.Result) int {
	totalWeight := 0
	items.ForEach(func(key, value gjson.Result) bool {
		weight := value.Get("grams").Int()
		quantity := value.Get("quantity").Int()
		totalWeight += int(weight * quantity)
		return true
	})
	return totalWeight
}

func ExtractOrderShopify(shopifyResponse string) Order {
	items := gjson.Get(shopifyResponse, "rate.items")

	return Order{
		Origin: Location{
			City:     gjson.Get(shopifyResponse, "rate.origin.city").String(),
			Id:       "",
			Province: gjson.Get(shopifyResponse, "rate.origin.province").String(),
		},
		Destination: Location{
			City:     gjson.Get(shopifyResponse, "rate.destination.city").String(),
			Id:       "",
			Province: gjson.Get(shopifyResponse, "rate.destination.province").String(),
		},
		Weight: CalculateWeight(items),
	}
}

func CreateShopifyResponse(rates []ShippingRate) ShopifyCarrierServiceResponse {
	return ShopifyCarrierServiceResponse{Rates: rates}
}
