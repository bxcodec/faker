package faker

import (
	"reflect"
	"strings"
	"testing"
)

func TestCreditCardType(t *testing.T) {
	ccType, err := GetPayment().CreditCardType(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	randCC := strings.ToLower(ccType.(string))

	if _, exist := creditCards[randCC]; !exist {
		t.Errorf("Expected from function creditCardType() : %s", randCC)
	}
}

func TestCreditCardNumber(t *testing.T) {
	_, err := GetPayment().CreditCardNumber(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
}

func TestSetPayment(t *testing.T) {
	SetPayment(Payment{})
}

func TestFakeCreditCardType(t *testing.T) {
	ccType := CCType()
	randCC := strings.ToLower(ccType)

	if _, exist := creditCards[randCC]; !exist {
		t.Errorf("Expected from function creditCardType() : %s", randCC)
	}
}

func TestFakeCreditCardNumber(t *testing.T) {
	ccNumber := CCNumber()

	if ccNumber == "" {
		t.Error("Expected Credit Card Number ")
	}
}
