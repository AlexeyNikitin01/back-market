package contract

type CartProduct struct {
	ID int `json:"product_id"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Discount int `json:"discount"`
	Amount int `json:"amount"`
}

type CartResponse struct {
	Products []CartProduct
	TotalAmount int `json:"totalAmount"` 
}

type CartRequest struct {
	Products []CartProductRequest `json:"products"`
}

type CartProductRequest struct {
	ID int `json:"id"`
	Quantity int `json:"quantity"`
}
