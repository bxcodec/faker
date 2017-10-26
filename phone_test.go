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
