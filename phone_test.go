package faker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/options"
)

func TestPhoneNumber(t *testing.T) {
	ph, err := GetPhoner().PhoneNumber(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if strings.Count(ph.(string), "-") != 2 {
		t.Error("Expected no more than two characters '-'")
	}
}

func TestTollFreePhoneNumber(t *testing.T) {
	ph, err := GetPhoner().TollFreePhoneNumber(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.HasPrefix(ph.(string), "(888)") && !strings.HasPrefix(ph.(string), "(777)") {
		t.Error("Expected character '(888)' or (777), in function TollFreePhoneNumber")
	}
}

func TestE164PhoneNumber(t *testing.T) {
	ph, err := GetPhoner().E164PhoneNumber(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.HasPrefix(ph.(string), "+") {
		t.Error("Expected character '(888)', in function TollFreePhoneNumber")
	}
}

func TestFakePhoneNumber(t *testing.T) {
	ph := Phonenumber(options.DefaultOption())
	if strings.Count(ph, "-") != 2 {
		t.Error("Expected no more than two characters '-'")
	}
}

func TestFakeTollFreePhoneNumber(t *testing.T) {
	ph := TollFreePhoneNumber(options.DefaultOption())
	if !strings.HasPrefix(ph, "(888)") && !strings.HasPrefix(ph, "(777)") {
		t.Error("Expected character '(888)' or (777), in function TollFreePhoneNumber")
	}
}

func TestFakeE164PhoneNumber(t *testing.T) {
	ph := E164PhoneNumber(options.DefaultOption())
	if !strings.HasPrefix(ph, "+") {
		t.Error("Expected character '(888)', in function TollFreePhoneNumber")
	}
}
