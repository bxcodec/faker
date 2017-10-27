package faker

import (
	"strings"
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	ph := getPhoner()
	if strings.Count(ph.PhoneNumber(), "-") != 2 {
		t.Error("Expected no more than two characters '-'")
	}
}

func TestTollFreePhoneNumber(t *testing.T) {
	ph := getPhoner()

	if !strings.HasPrefix(ph.TollFreePhoneNumber(), "(888)") && !strings.HasPrefix(ph.TollFreePhoneNumber(), "(777)") {
		t.Error("Expected character '(888)' or (777), in function TollFreePhoneNumber")
	}
}

func TestE164PhoneNumber(t *testing.T) {
	ph := getPhoner()
	if !strings.HasPrefix(ph.E164PhoneNumber(), "+") {
		t.Error("Expected character '(888)', in function TollFreePhoneNumber")
	}
}

func TestSetPhoner(t *testing.T) {
	SetPhoner(Phone{})
}
