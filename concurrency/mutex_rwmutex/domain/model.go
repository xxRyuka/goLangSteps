package domain

type OrderType string

const (
	Sell OrderType = "Sell"
	Buy  OrderType = "Buy"
)

type Order struct {
	Symbol string
	Amount float64
	Type   OrderType // sell or buy

}
