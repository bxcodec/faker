package faker

import (
	"github.com/bxcodec/faker/support/slice"
	"testing"
)

func TestSetDowser(t *testing.T) {
	SetDowser(Person{})
}

func TestTitleMale(t *testing.T) {
	p := getPerson()
	if !slice.Contains(titlesMale, p.TitleMale()) {
		t.Error("Expected value from variable titleMales in function TitleMale")
	}
}

func TestTitleFemale(t *testing.T) {
	p := getPerson()
	if !slice.Contains(titlesFemales, p.TitleFeMale()) {
		t.Error("Expected value from variable titleFemales in function TitleFeMale")
	}
}

func TestFirstNameMale(t *testing.T) {
	p := getPerson()
	if !slice.Contains(firstNamesMale, p.FirstNameMale()) {
		t.Error("Expected value from variable firstNamesMale in function FirstNameMale")
	}
}

func TestFirstNameFemale(t *testing.T) {
	p := getPerson()
	if !slice.Contains(firstNamesFemale, p.FirstNameFemale()) {
		t.Error("Expected value from variable firstNamesFemale in function FirstNameFemale")
	}
}

func TestLastName(t *testing.T) {
	p := getPerson()
	if !slice.Contains(lastNames, p.LastName()) {
		t.Error("Expected value from variable lastNames in function LastName")
	}
}
