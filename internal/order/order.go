package order

import "time"

type OrderType string

const (
	Buy  OrderType = "BUY"
	Sell OrderType = "SELL"
)

type Order struct {
	ID        string
	UserID    string
	Symbol    string
	Quantity  int
	Price     float64
	Type      OrderType
	Timestamp time.Time
}
