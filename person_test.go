package faker

import (
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestSetDowser(t *testing.T) {
	SetDowser(Person{})
}

func TestTitleMale(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(titlesMale, p.TitleMale()) {
		t.Error("Expected value from variable titleMales in function TitleMale")
	}
}

func TestTitleFemale(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(titlesFemale, p.TitleFeMale()) {
		t.Error("Expected value from variable titleFemales in function TitleFeMale")
	}
}

func TestFirstNameMale(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(firstNamesMale, p.FirstNameMale()) {
		t.Error("Expected value from variable firstNamesMale in function FirstNameMale")
	}
}

func TestFirstNameFemale(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(firstNamesFemale, p.FirstNameFemale()) {
		t.Error("Expected value from variable firstNamesFemale in function FirstNameFemale")
	}
}

func TestFirstName(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(firstNames, p.FirstName()) {
		t.Error("Expected value from either firstNamesMale or firstNamesFemale in function FirstName")
	}
}

func TestLastName(t *testing.T) {
	p := GetPerson()
	if !slice.Contains(lastNames, p.LastName()) {
		t.Error("Expected value from variable lastNames in function LastName")
	}
}

func TestNameMale(t *testing.T) {
	p := GetPerson()
	randNameFlag = 51
	if p.Name() == "" {
		t.Error("Expected from function name string get empty string")
	}
}
func TestNameFemale(t *testing.T) {
	p := GetPerson()
	randNameFlag = 20
	if p.Name() == "" {
		t.Error("Expected from function name string get empty string")
	}
}
