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
	GetPayment().CreditCardNumber(reflect.Value{})
}

func TestSetPayment(t *testing.T) {
	SetPayment(Payment{})
}
