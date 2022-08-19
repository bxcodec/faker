package faker

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestDigit(t *testing.T) {
	p := GetIdentifier()
	uuid, err := p.Digit(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if match, err := regexp.Match("^[a-zA-Z0-9]{32}$", []byte(uuid.(string))); !match || err != nil {
		t.Errorf("Could not match the UUID format, err: %+v, match: %+v", err, match)
	}
}

func TestHyphenated(t *testing.T) {
	p := GetIdentifier()
	uuid, err := p.Hyphenated(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	exp := "[a-zA-Z 0-9]"
	pattern := fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$", exp, exp, exp, exp, exp)
	if match, err := regexp.Match(pattern, []byte(uuid.(string))); !match || err != nil {
		t.Errorf("Could not match the UUID hyphenated format, err: %+v, match: %+v", err, match)
	}

}

func TestGetIdentifier(t *testing.T) {
	identifier := GetIdentifier()
	if identifier == nil {
		t.Fatalf("TestGetIdentifier failed because identifier was nil")
	}
}

func TestFakeDigit(t *testing.T) {
	uuid := UUIDDigit()
	if match, err := regexp.Match("^[a-zA-Z0-9]{32}$", []byte(uuid)); !match || err != nil {
		t.Errorf("Could not match the UUID format, err: %+v, match: %+v", err, match)
	}
}

func TestFakeHyphenated(t *testing.T) {
	uuid := UUIDHyphenated()
	exp := "[a-zA-Z 0-9]"
	pattern := fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$", exp, exp, exp, exp, exp)
	if match, err := regexp.Match(pattern, []byte(uuid)); !match || err != nil {
		t.Errorf("Could not match the UUID hyphenated format, err: %+v, match: %+v", err, match)
	}
}
