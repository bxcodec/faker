package faker

import (
	"testing"

	"github.com/bxcodec/faker/v4/options"
)

func TestGetLongitude(t *testing.T) {
	long := Longitude(options.DefaultOption())
	if long > 180 || long < -180 {
		t.Error("function Longitude need return a valid longitude")
	}
}

func TestGetLatitude(t *testing.T) {
	lat := Latitude(options.DefaultOption())
	if lat > 90 || lat < -90 {
		t.Error("function Latitude need return a valid longitude")
	}
}
