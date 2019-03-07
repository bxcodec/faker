package faker

import (
	"testing"
)

func TestFakeData(t *testing.T) {
	SetAddress(Address{})
}

func TestGetLongitude(t *testing.T) {
	long := Longitude()
	if long > 180 || long < -180 {
		t.Error("function Longitude need return a valid longitude")
	}
}

func TestGetLatitude(t *testing.T) {
	lat := Latitude()
	if lat > 90 || lat < -90 {
		t.Error("function Latitude need return a valid longitude")
	}
}
