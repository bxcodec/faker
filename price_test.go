package faker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestSetPrice(t *testing.T) {
	SetPrice(Price{})
}

func TestCurrency(t *testing.T) {
	p := GetPrice()
	fmt.Println(p.Currency())
	if !slice.Contains(currencies, p.Currency()) {
		t.Error("Expected a currency code from currencies")
	}
}

func TestAmountWithCurrency(t *testing.T) {
	p := GetPrice()
	fmt.Println(p.AmountWithCurrency())
	if !strings.Contains(p.AmountWithCurrency(), " ") {
		t.Error("Expected Price currency followed by a space and it's ammount")
	}
}
