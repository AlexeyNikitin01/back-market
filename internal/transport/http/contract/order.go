package contract

import "time"

type OrderProduct struct {
	ProductID int `json:"id"`
	Quantity  int `json:"qantity"`
	Price     int `json:"price"`
	Discount  int `json:"discount"`
	Amount    int `json:"amount"`
}

type OrderResponse struct {
	ID          int            `json:"id"`
	Products    []OrderProduct `json:"products"`
	CreateAt    time.Time      `json:"create_at"`
	TotalAmount int            `json:"totalAmount"`
	Paid 		bool		   `json:"paid"`
}

type HistoryResponse struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	ID       int       `json:"id"`
	CreateAt time.Time `json:"createdAt"`
	Paid     bool      `json:"paid"`
}

type OrderRequest struct {
	Address string `json:"address"`
}
