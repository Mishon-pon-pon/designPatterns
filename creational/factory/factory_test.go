package factory

import (
	"strings"
	"testing"
)

func TestPaymentCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}
	msg := payment.Pay(10.12)
	if !strings.Contains(msg, "paid using Cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestPaymentDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment of type 'DebitCard' must exist")
	}

	msg := payment.Pay(10.3)
	if !strings.Contains(msg, "paid using DebitCard") {
		t.Error("The DebitCard payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
