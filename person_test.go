package faker

import (
	"reflect"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestSetDowser(t *testing.T) {
	SetDowser(Person{})
}

func TestTitleMale(t *testing.T) {
	male, err := GetPerson().TitleMale(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(titlesMale, male.(string)) {
		t.Error("Expected value from variable titleMales in function TitleMale")
	}
}

func TestTitleFemale(t *testing.T) {
	female, err := GetPerson().TitleFeMale(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(titlesFemale, female.(string)) {
		t.Error("Expected value from variable titleFemales in function TitleFeMale")
	}
}

func TestFirstNameMale(t *testing.T) {
	firstName, err := GetPerson().FirstNameMale(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(firstNamesMale, firstName.(string)) {
		t.Error("Expected value from variable firstNamesMale in function FirstNameMale")
	}
}

func TestFirstNameFemale(t *testing.T) {
	firstName, err := GetPerson().FirstNameFemale(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(firstNamesFemale, firstName.(string)) {
		t.Error("Expected value from variable firstNamesFemale in function FirstNameFemale")
	}
}

func TestFirstName(t *testing.T) {
	firstname, err := GetPerson().FirstName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(firstNames, firstname.(string)) {
		t.Error("Expected value from either firstNamesMale or firstNamesFemale in function FirstName")
	}
}

func TestLastName(t *testing.T) {
	lastname, err := GetPerson().LastName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(lastNames, lastname.(string)) {
		t.Error("Expected value from variable lastNames in function LastName")
	}
}

func TestNameMale(t *testing.T) {
	name, err := GetPerson().Name(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	randNameFlag = 51
	if name.(string) == "" {
		t.Error("Expected from function name string get empty string")
	}
}
func TestNameFemale(t *testing.T) {
	name, err := GetPerson().Name(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	randNameFlag = 20
	if name.(string) == "" {
		t.Error("Expected from function name string get empty string")
	}
}
