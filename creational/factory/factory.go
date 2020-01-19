package factory

import "errors"

type PaymentMethod interface {
	Pay(n float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(i int) (PaymentMethod, error) {
	return nil, errors.New("Not implemented yet")
}

type CashPM struct{}
type DebitCardPM struct{}

func (CashPM) Pay(n float32) string {
	return ""
}

func (DebitCardPM) Pay(n float32) string {
	return ""
}
