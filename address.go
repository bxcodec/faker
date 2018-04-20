package faker

import (
	"math/rand"
	"reflect"
)

var address Addresser

// Constructor
func GetAddress() Addresser {
	mu.Lock()
	defer mu.Unlock()

	if address == nil {
		address = &Address{}
	}
	return address
}

// this set custom Address
func SetAddress(net Addresser) {
	address = net
}

type Addresser interface {
	Latitude(v reflect.Value) error
	Longitude(v reflect.Value) error
}

type Address struct{}

func (i Address) Latitude(v reflect.Value) error {
	kind := v.Kind()
	val := (rand.Float32() * 180) - 90
	if kind == reflect.Float32 {
		v.Set(reflect.ValueOf(val))
		return nil
	}
	v.Set(reflect.ValueOf(float64(val)))

	return nil
}

func (i Address) Longitude(v reflect.Value) error {
	kind := v.Kind()
	val := (rand.Float32() * 360) - 180

	if kind == reflect.Float32 {
		v.Set(reflect.ValueOf(val))
		return nil
	}
	v.Set(reflect.ValueOf(float64(val)))
	return nil
}
