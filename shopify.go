package columba

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

type Order struct {
	Origin      Location
	Destination Location
	Items       []Item
	Currency    string
	Locale      string
}

type ShippingRate struct {
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	TotalPrice  string `json:"total_price"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
}

type CarrierServiceQuery struct {
	Rate Order
}

type CarrierServiceResponse struct {
	Rates []ShippingRate `json:"rates"`
}
