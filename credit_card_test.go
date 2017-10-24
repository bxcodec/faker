package faker

import (
	"testing"
	"strings"
)

func TestCreditCardType(t *testing.T)  {
	randCC := strings.ToLower(creditCardType())
	if _, exist :=  creditCards[randCC]; !exist {
		t.Errorf("Expected from function creditCardType() : %s", randCC)
	}
}

func TestCreditCardNumIfArgumentEmpty(t *testing.T)  {
	cacheCreditCard = ""
	creditCardNum("")
}

func TestCreditCardNumIfArgumentExist(t *testing.T)  {
	creditCardNum(strings.ToLower(creditCardType()))
}

func TestCreditCardNumIfExistCache(t *testing.T)  {
	cacheCreditCard = strings.ToLower(creditCardType())
	creditCardNum("")
}