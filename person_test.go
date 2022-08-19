package faker

import (
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v4/pkg/slice"
)

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

func TestFakeTitleMale(t *testing.T) {
	male := TitleMale()
	if !slice.Contains(titlesMale, male) {
		t.Error("Expected value from variable titleMales in function TitleMale")
	}
}

func TestFakeTitleFemale(t *testing.T) {
	female := TitleFemale()
	if !slice.Contains(titlesFemale, female) {
		t.Error("Expected value from variable titleFemales in function TitleFeMale")
	}
}

func TestFakeFirstNameMale(t *testing.T) {
	firstName := FirstNameMale()
	if !slice.Contains(firstNamesMale, firstName) {
		t.Error("Expected value from variable firstNamesMale in function FirstNameMale")
	}
}

func TestFakeFirstNameFemale(t *testing.T) {
	firstName := FirstNameFemale()
	if !slice.Contains(firstNamesFemale, firstName) {
		t.Error("Expected value from variable firstNamesFemale in function FirstNameFemale")
	}
}

func TestFakeFirstName(t *testing.T) {
	firstname := FirstName()
	if !slice.Contains(firstNames, firstname) {
		t.Error("Expected value from either firstNamesMale or firstNamesFemale in function FirstName")
	}
}

func TestFakeLastName(t *testing.T) {
	lastname := LastName()
	if !slice.Contains(lastNames, lastname) {
		t.Error("Expected value from variable lastNames in function LastName")
	}
}

func TestFakeNameMale(t *testing.T) {
	name := Name()
	randNameFlag = 51
	if name == "" {
		t.Error("Expected from function name string get empty string")
	}
}
func TestFakeNameFemale(t *testing.T) {
	name := Name()
	randNameFlag = 20
	if name == "" {
		t.Error("Expected from function name string get empty string")
	}
	t.Log(name)
}

func TestFakeGender(t *testing.T) {
	gender, err := GetPerson().Gender(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(genders, gender.(string)) {
		t.Error("Expected value from variable genders in function Gender")
	}
}

func TestFakeGenderPublicFunction(t *testing.T) {
	gender := Gender()
	if !slice.Contains(genders, gender) {
		t.Error("Expected value from variable genders in function Gender")
	}
}

func TestChineseFirstName(t *testing.T) {
	firstname, err := GetPerson().ChineseFirstName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	t.Log(firstname)
	if !slice.Contains(chineseFirstNames, firstname.(string)) {
		t.Error("Expected value from either chineseFirstNames in function ChineseFirstName")
	}
}

func TestChineseLastName(t *testing.T) {
	firstname, err := GetPerson().ChineseLastName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	t.Log(firstname)
	if !slice.Contains(chineseLastNames, firstname.(string)) {
		t.Error("Expected value from either chineseLastNames in function ChineseLastName")
	}
}

func TestChineseName(t *testing.T) {
	firstname, err := GetPerson().ChineseName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	t.Log(firstname)
}
