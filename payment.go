package faker

import (
	"math/rand"
	"reflect"
	"strconv"
)

const (
	numberBytes = "0123456789"
)

// creditCard struct
type creditCard struct {
	ccType   string
	length   int
	prefixes []int
}

var creditCards = map[string]creditCard{
	"visa":             {"VISA", 16, []int{4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4}},
	"mastercard":       {"MasterCard", 16, []int{51, 52, 53, 54, 55}},
	"american express": {"American Express", 15, []int{34, 37}},
	"discover":         {"Discover", 16, []int{6011}},
	"VISA":             {"VISA", 16, []int{4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4}},
	"MasterCard":       {"MasterCard", 16, []int{51, 52, 53, 54, 55}},
	"American Express": {"American Express", 15, []int{34, 37}},
	"Discover":         {"Discover", 16, []int{6011}},
}

var pay Render

var cacheCreditCard string

// GetPayment returns a new Render interface of Payment struct
func GetPayment() Render {
	mu.Lock()
	defer mu.Unlock()

	if pay == nil {
		pay = &Payment{}
	}
	return pay
}

// SetPayment set custom Network
func SetPayment(p Render) {
	pay = p
}

// Render contains Whole Random Credit Card Generators with their types
type Render interface {
	CreditCardType(v reflect.Value) (interface{}, error)
	CreditCardNumber(v reflect.Value) (interface{}, error)
}

// Payment struct
type Payment struct{}

func (p Payment) cctype() string {
	n := len(creditCards)
	if cacheCreditCard != "" {
		return cacheCreditCard
	}
	var ccTypes []string

	for _, cc := range creditCards {
		ccTypes = append(ccTypes, cc.ccType)
	}
	cacheCreditCard = ccTypes[rand.Intn(n)]
	return cacheCreditCard
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func (p Payment) CreditCardType(v reflect.Value) (interface{}, error) {
	return p.cctype(), nil
}

func (p Payment) ccnumber() string {
	ccType := p.cctype()
	cacheCreditCard = ccType
	card := creditCards[ccType]
	prefix := strconv.Itoa(card.prefixes[rand.Intn(len(card.prefixes))])

	num := prefix
	digit := randomStringNumber(card.length - len(prefix))

	num += digit
	return num
}

// CreditCardNumber generated credit card number according to the card number rules
func (p Payment) CreditCardNumber(v reflect.Value) (interface{}, error) {
	return p.ccnumber(), nil
}
