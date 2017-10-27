package faker

import (
	"testing"
	"reflect"
	"time"
	"log"
)

func TestSetDateTimer(t *testing.T) {
	SetDateTimer(DateTime{})
}

func TestUnixTimeValueValid(t *testing.T) {
	d := getDateTimer()
	var ref = struct {
		some int64
	}{
		some: 1212,
	}
	d.UnixTime(reflect.ValueOf(&ref.some).Elem())

	if time.Now().Unix() <= ref.some {
		t.Error("UnixTime should return time <= now")
	}
}
func TestUnixTimeValueNotValid(t *testing.T) {
	d := getDateTimer()
	var ref = struct {
		some int
	}{
		some: 1212,
	}
	d.UnixTime(reflect.ValueOf(&ref.some).Elem())
	log.Println(ref.some)
	if ref.some != 0{
		t.Errorf("UnixTime should return 0, get : %v", ref.some)
	}
}