package domain

type PaymentProccessor interface {
	Process(amount float64) error
}
