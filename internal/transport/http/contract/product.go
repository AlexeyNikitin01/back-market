package contract

type ProductsRequest []ProductRequest

type ProductRequest struct {
	Name             string  `json:"name"`
	Category 		 string  `json:"category"`
	Price            float64 `json:"price"`
	PremiumDiscount  float64 `json:"premiumDiscount,omitempty"`
	CategoryDiscount float64 `json:"categoryDiscount,omitempty"`
}

type ProductsResponse struct {
	Products   []ProductResponse `json:"products,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

type ProductResponse struct {
	Id       int64   `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Category string  `json:"category,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}

type ProductsUpdateRequest []ProductUpdateRequest

type ProductUpdateRequest struct {
	Id				 int     `json:"product_id"`
	Name             string  `json:"name"`
	Category 		 string  `json:"category"`
	Price            float64 `json:"price"`
	PremiumDiscount  float64 `json:"premiumDiscount,omitempty"`
	CategoryDiscount float64 `json:"categoryDiscount,omitempty"`
}