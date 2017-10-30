package faker

import (
	"strings"
	"testing"
)

func TestCreditCardType(t *testing.T) {
	randCC := strings.ToLower(GetPayment().CreditCardType())
	if _, exist := creditCards[randCC]; !exist {
		t.Errorf("Expected from function creditCardType() : %s", randCC)
	}
}

func TestCreditCardNumber(t *testing.T) {
	GetPayment().CreditCardNumber()
}

func TestSetPayment(t *testing.T)  {
	SetPayment(Payment{})
}