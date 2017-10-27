package faker

import (
	"log"
	"reflect"
	"testing"
	"time"
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
	if ref.some != 0 {
		t.Errorf("UnixTime should return 0, get : %v", ref.some)
	}
}

func TestDate(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(BaseDate, d.Date())

	if err != nil {
		t.Error("function Date need return valid value")
	}
}

func TestTime(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(Time, d.Time())

	if err != nil {
		t.Error("function Time need return valid value")
	}
}

func TestMonth(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(Mounth, d.Month())
	if err != nil {
		t.Error("function Month need return valid month")
	}
}

func TestYear(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(Year, d.Year())
	if err != nil {
		t.Error("function Year need return valid year")
	}
}