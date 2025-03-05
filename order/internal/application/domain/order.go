package domain

import "time"

type ProductItem struct {
	ProductId string  `json:"product_id"`
	UnitPrice float32 `json:"unit_price"`
	Quantity  int32   `json:"quantity"`
}
type Order struct {
	ID         string        `json:"id"`
	CustomerId string        `json:"customer_id"`
	Status     string        `json:"status"`
	OrderItems []ProductItem `json:"order_items"`
	CreatedAt  int64         `json:"created_at"`
}

func NewOrder(customerId string, items []ProductItem) *Order {
	return &Order{
		CustomerId: customerId,
		Status:     "pending",
		OrderItems: items,
		CreatedAt:  time.Now().Unix(),
	}
}
