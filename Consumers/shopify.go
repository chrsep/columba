package Consumers

import (
	"columba"
	"github.com/tidwall/gjson"
)

type Item struct {
	Name               string
	Sku                string
	Quantity           string
	Grams              string
	Price              string
	Vendor             string
	RequiresShipping   bool
	Taxable            bool
	FulfillmentService string
	Properties         string
	ProductId          int
	VariantId          int
}

type Location struct {
	Country     string
	PostalCode  string
	Province    string
	City        string
	Name        string
	Address1    string
	Address2    string
	Address3    string
	Phone       string
	Fax         string
	Email       string
	AddressType string
	CompanyName string
}

type CarrierServiceResponse struct {
	Rates []columba.ShippingRate `json:"rates"`
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

func ExtractOrderShopify(shopifyResponse string) columba.Order {
	items := gjson.Get(shopifyResponse, "rate.items")

	return columba.Order{
		Origin: columba.Location{
			City:     gjson.Get(shopifyResponse, "rate.origin.city").String(),
			Id:       "",
			Province: gjson.Get(shopifyResponse, "rate.origin.province").String(),
		},
		Destination: columba.Location{
			City:     gjson.Get(shopifyResponse, "rate.origin.city").String(),
			Id:       "",
			Province: gjson.Get(shopifyResponse, "rate.origin.province").String(),
		},
		Weight: CalculateWeight(items),
	}
}
