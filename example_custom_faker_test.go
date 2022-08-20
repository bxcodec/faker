package faker_test

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker/v4"
)

// Gondoruwo ...
type Gondoruwo struct {
	Name       string
	Locatadata int
}

// custom type that aliases over slice of byte
type CustomUUID []byte

// Sample ...
type Sample struct {
	ID        int64      `faker:"customIdFaker"`
	Gondoruwo Gondoruwo  `faker:"gondoruwo"`
	Danger    string     `faker:"danger"`
	UUID      CustomUUID `faker:"customUUID"`
}

// CustomGenerator ...
func CustomGenerator() {
	_ = faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		return int64(43), nil
	})
	_ = faker.AddProvider("danger", func(v reflect.Value) (interface{}, error) {
		return "danger-ranger", nil
	})

	_ = faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
		obj := Gondoruwo{
			Name:       "Power",
			Locatadata: 324,
		}
		return obj, nil
	})

	_ = faker.AddProvider("customUUID", func(v reflect.Value) (interface{}, error) {
		s := CustomUUID{
			0, 8, 7, 2, 3,
		}
		return s, nil
	})
}

// You can also add your own generator function to your own defined tags.
func Example_customFaker() {
	CustomGenerator()
	var sample Sample
	_ = faker.FakeData(&sample)
	fmt.Printf("%+v", sample)
	// Output:
	// {ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger UUID:[0 8 7 2 3]}
}
