package faker_test

import (
	"fmt"
	"reflect"

	"github.com/bxcodec/faker/v3"
)

// Gondoruwo ...
type Gondoruwo struct {
	Name       string
	Locatadata int
}

// Sample ...
type Sample struct {
	ID        int64     `faker:"customIdFaker"`
	Gondoruwo Gondoruwo `faker:"gondoruwo"`
	Danger    string    `faker:"danger"`
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
}

// You can also add your own generator function to your own defined tags.
func Example_customFaker() {
	CustomGenerator()
	var sample Sample
	_ = faker.FakeData(&sample)
	fmt.Printf("%+v", sample)
	// Output:
	// {ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger}
}
