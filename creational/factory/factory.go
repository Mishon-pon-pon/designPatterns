package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(n float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (CashPM) Pay(n float32) string {
	return fmt.Sprintf("%0.2f paid using Cash\n", n)
}

func (DebitCardPM) Pay(n float32) string {
	return fmt.Sprintf("%0.2f paid using DebitCard\n", n)
}
