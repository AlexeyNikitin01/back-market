package contract

type PayRequest struct {
	PaymentType string `json:"paymentType"`
	Amount      int    `json:"amount"`
	OrderID     int    `json:"order_id,omitempty"`
}
