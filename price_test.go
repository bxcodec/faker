package faker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestSetPrice(t *testing.T) {
	SetPrice(Price{})
}

func TestCurrency(t *testing.T) {
	p, err := GetPrice().Currency(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(currencies, p.(string)) {
		t.Error("Expected a currency code from currencies")
	}
}

func TestAmountWithCurrency(t *testing.T) {
	p, err := GetPrice().AmountWithCurrency(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}

	if !strings.Contains(p.(string), " ") {
		t.Error("Expected Price currency followed by a space and it's ammount")
	}
}
