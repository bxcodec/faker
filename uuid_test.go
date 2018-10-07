package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDigit(t *testing.T) {
	p := GetIdentifier()
	uuid := p.Digit()
	if match, err := regexp.Match("^[a-zA-Z0-9]{32}$", []byte(uuid)); !match || err != nil {
		t.Errorf("Could not match the UUID format, err: %+v, match: %+v", err, match)
	}
}

func TestHyphenated(t *testing.T) {
	p := GetIdentifier()
	uuid := p.Hyphenated()
	exp := "[a-zA-Z 0-9]"
	pattern := fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$", exp, exp, exp, exp, exp)
	if match, err := regexp.Match(pattern, []byte(uuid)); !match || err != nil {
		t.Errorf("Could not match the UUID hyphenated format, err: %+v, match: %+v", err, match)
	}

}
