package faker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/options"
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

func TestFakeCreditCardType(t *testing.T) {
	ccType := CCType(options.DefaultOption())
	randCC := strings.ToLower(ccType)

	if _, exist := creditCards[randCC]; !exist {
		t.Errorf("Expected from function creditCardType() : %s", randCC)
	}
}

func TestFakeCreditCardNumber(t *testing.T) {
	ccNumber := CCNumber(options.DefaultOption())

	if ccNumber == "" {
		t.Error("Expected Credit Card Number ")
	}
}
