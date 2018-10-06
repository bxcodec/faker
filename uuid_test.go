package faker

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

const (
	Tag = `faker:"uuid_digit"`
)

type MockUUID struct {
	UUID           string `faker:"uuid_digit"`
	UUIDHyphenated string `faker:"uuid_hyphenated"`
}

/*func TestSetUUID(t *testing.T) {
	SetUUID(NogetAndet{})
}*/

func TestDigit(t *testing.T) {
	p := GetUUID()
	mock := MockUUID{}
	val := reflect.ValueOf(&mock.UUID)
	err := p.Digit(val.Elem())
	if err != nil {
		t.Error(err)
	}
	if match, err := regexp.Match("^[a-zA-Z0-9]{32}$", []byte(val.Elem().String())); !match || err != nil {
		t.Errorf("Could not match the UUID format, err: %+v, match: %+v", err, match)
	}
}

func TestHyphenated(t *testing.T) {
	p := GetUUID()
	mock := MockUUID{}
	val := reflect.ValueOf(&mock.UUIDHyphenated)
	err := p.Hyphenated(val.Elem())
	if err != nil {
		t.Error(err)
	}
	exp := "[a-zA-Z 0-9]"
	pattern := fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$", exp, exp, exp, exp, exp)
	if match, err := regexp.Match(pattern, []byte(val.Elem().String())); !match || err != nil {
		t.Errorf("Could not match the UUID hyphenated format, err: %+v, match: %+v", err, match)
	}

}
