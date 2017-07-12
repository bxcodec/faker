package faker

import (
	"math/rand"
	"strconv"
	"strings"
)

type creditCard struct {
	ccType   string
	length   int
	prefixes []int
}

const (
	numberBytes = "0123456789"
)

var creditCards = map[string]creditCard{
	"visa":             {"VISA", 16, []int{4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4}},
	"mastercard":       {"MasterCard", 16, []int{51, 52, 53, 54, 55}},
	"american express": {"American Express", 15, []int{34, 37}},
	"discover":         {"Discover", 16, []int{6011}},
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
var cacheCreditCard string

func creditCardType() string {
	n := len(creditCards)
	r := rand.New(src)
	if cacheCreditCard != "" {
		return cacheCreditCard
	}
	var ccTypes []string

	for _, cc := range creditCards {
		ccTypes = append(ccTypes, cc.ccType)
	}
	cacheCreditCard = ccTypes[r.Intn(n)]
	return cacheCreditCard
}

// CreditCardNum generated credit card number according to the card number rules
func creditCardNum(ccType string) string {
	r := rand.New(src)
	if ccType != "" {
		ccType = strings.ToLower(ccType)
		cacheCreditCard = ccType
	} else if cacheCreditCard != "" {
		ccType = strings.ToLower(cacheCreditCard)
	} else {
		ccType = strings.ToLower(creditCardType())
		cacheCreditCard = ccType
	}
	card := creditCards[ccType]
	prefix := strconv.Itoa(card.prefixes[r.Intn(len(card.prefixes))])

	num := prefix
	digit := randomStringNumber(card.length - len(prefix))

	num += digit
	return num
}

func randomStringNumber(n int) string {

	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
