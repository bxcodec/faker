package faker

import (
	"fmt"
	"github.com/bxcodec/faker/support/slice"
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

func TestMonthName(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(Month, d.MonthName())
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

func TestDayOfWeek(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(Day, d.DayOfWeek())
	if err != nil {
		t.Error("function DayOfWeek need return valid day")
	}
}

func TestDayOfMonth(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(DayOfMonth, d.DayOfMonth())
	if err != nil {
		t.Error("function DayOfMonth need return valid digit")
	}
}

func TestTimestamp(t *testing.T) {
	d := getDateTimer()
	_, err := time.Parse(fmt.Sprintf("%s %s", BaseDate, Time), d.Timestamp())
	if err != nil {
		t.Error("function Timestamp need return valid timestamp format")
	}
}

func TestCentury(t *testing.T) {
	d := getDateTimer()

	if !slice.Contains(century, d.Century()) {
		t.Error("Expected century from functuon Century")
	}
}

func TestTimeZone(t *testing.T) {
	d := getDateTimer()
	if !slice.Contains(timezones, d.TimeZone()) {
		t.Error("Expected timezone from variable timezones")
	}

}
