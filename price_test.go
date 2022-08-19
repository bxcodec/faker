package faker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/pkg/slice"
)

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
		t.Error("Expected Price currency followed by a space and it's amount")
	}
}

func TestFakeCurrency(t *testing.T) {
	p := Currency()
	if !slice.Contains(currencies, p) {
		t.Error("Expected a currency code from currencies")
	}
}

func TestFakeAmountWithCurrency(t *testing.T) {
	p := AmountWithCurrency()

	if !strings.Contains(p, " ") {
		t.Error("Expected Price currency followed by a space and it's amount")
	}

	arrCurrency := strings.Split(p, " ")

	if !slice.Contains(currencies, arrCurrency[0]) {
		t.Error("Expected a currency code from currencies")
	}
}
