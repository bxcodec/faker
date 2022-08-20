package faker

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4/pkg/slice"
)

func TestUnixTimeValueValid(t *testing.T) {
	d := GetDateTimer()
	var ref = struct {
		some int64
	}{
		some: 1212,
	}
	_, err := d.UnixTime(reflect.ValueOf(&ref.some).Elem())
	if err != nil {
		t.Error("function Date need return valid value")
	}

	if time.Now().Unix() <= ref.some {
		t.Error("UnixTime should return time <= now")
	}
}

func TestDate(t *testing.T) {
	d := GetDateTimer()
	date, err := d.Date(reflect.Value{})
	if err != nil {
		t.Error("function Date need return valid value")
	}
	_, err = time.Parse(BaseDateFormat, date.(string))

	if err != nil {
		t.Error("function Date need return valid value")
	}
}

func TestTime(t *testing.T) {
	d := GetDateTimer()
	tm, err := d.Time(reflect.Value{})
	if err != nil {
		t.Error("function Time need return valid value")
	}
	_, err = time.Parse(TimeFormat, tm.(string))
	if err != nil {
		t.Error("function Time need return valid value")
	}
}

func TestMonthName(t *testing.T) {
	d := GetDateTimer()
	mt, err := d.MonthName(reflect.Value{})
	if err != nil {
		t.Error("function Month need return valid month")
	}
	_, err = time.Parse(MonthFormat, mt.(string))
	if err != nil {
		t.Error("function Month need return valid month")
	}
}

func TestYear(t *testing.T) {
	d := GetDateTimer()
	year, err := d.Year(reflect.Value{})
	if err != nil {
		t.Error("function Year need return valid year")
	}
	_, err = time.Parse(YearFormat, year.(string))
	if err != nil {
		t.Error("function Year need return valid year")
	}
}

func TestDayOfWeek(t *testing.T) {
	d := GetDateTimer()
	week, err := d.DayOfWeek(reflect.Value{})
	if err != nil {
		t.Error("function DayOfWeek need return valid day")
	}
	_, err = time.Parse(DayFormat, week.(string))
	if err != nil {
		t.Error("function DayOfWeek need return valid day")
	}
}

func TestDayOfWeekReturnsDifferentValues(t *testing.T) {
	dayMap := make(map[string]struct{})
	iterations := 5 // sufficiently large to assure we don't randomly get the same value again
	for i := 0; i < iterations; i++ {
		day, err := GetDateTimer().DayOfWeek(reflect.Value{})
		if err != nil {
			t.Error("function DayOfWeek need return valid day")
		}
		if _, ok := dayMap[day.(string)]; ok {
			i--
			continue
		}
		dayMap[day.(string)] = struct{}{}
		t.Log(day) // Will print random and different day 5 times.
	}

	if len(dayMap) < 1 {
		t.Error("function need return at least one day item")
	}
}

func TestDayOfMonth(t *testing.T) {
	d := GetDateTimer()
	mt, err := d.DayOfMonth(reflect.Value{})
	if err != nil {
		t.Error("function DayOfMonth need return valid digit")
	}
	_, err = time.Parse(DayOfMonthFormat, mt.(string))
	if err != nil {
		t.Error("function DayOfMonth need return valid digit")
	}
}

func TestTimestamp(t *testing.T) {
	d := GetDateTimer()
	tstmp, err := d.Timestamp(reflect.Value{})
	if err != nil {
		t.Error("function Timestamp need return valid timestamp format")
	}
	_, err = time.Parse(fmt.Sprintf("%s %s", BaseDateFormat, TimeFormat), tstmp.(string))
	if err != nil {
		t.Error("function Timestamp need return valid timestamp format")
	}
}

func TestCentury(t *testing.T) {
	d := GetDateTimer()
	centry, err := d.Century(reflect.Value{})
	if err != nil {
		t.Error("Expected century from functuon Century")
	}
	if !slice.Contains(century, centry.(string)) {
		t.Error("Expected century from functuon Century")
	}
}

func TestTimeZone(t *testing.T) {
	d := GetDateTimer()
	tz, err := d.TimeZone(reflect.Value{})
	if err != nil {
		t.Error("Expected timezone from variable timezones")
	}
	if !slice.Contains(timezones, tz.(string)) {
		t.Error("Expected timezone from variable timezones")
	}
}

func TestTimePeriod(t *testing.T) {
	d := GetDateTimer()
	periode, err := d.TimePeriod(reflect.Value{})
	if err != nil {
		t.Error("function TimePeriod need return valid period")
	}
	_, err = time.Parse(TimePeriodFormat, periode.(string))
	if err != nil {
		t.Error("function TimePeriod need return valid period")
	}
}

func TestFakeUnixTime(t *testing.T) {
	unixTime := UnixTime()

	if time.Now().Unix() <= unixTime {
		t.Error("UnixTime should return time <= now")
	}
}

func TestFakeDate(t *testing.T) {
	date := Date()

	_, err := time.Parse(BaseDateFormat, date)
	if err != nil {
		t.Error("function Date need return valid value")
	}
}

func TestFakeTime(t *testing.T) {
	tm := TimeString()
	_, err := time.Parse(TimeFormat, tm)
	if err != nil {
		t.Error("function Time need return valid value")
	}
}

func TestFakeMonthName(t *testing.T) {
	mt := MonthName()
	_, err := time.Parse(MonthFormat, mt)
	if err != nil {
		t.Error("function Month need return valid month")
	}
}

func TestFakeYear(t *testing.T) {
	year := YearString()
	_, err := time.Parse(YearFormat, year)
	if err != nil {
		t.Error("function Year need return valid year")
	}
}

func TestFakeDayOfWeek(t *testing.T) {
	week := DayOfWeek()
	_, err := time.Parse(DayFormat, week)
	if err != nil {
		t.Error("function DayOfWeek need return valid day")
	}
}

func TestFakeDayOfMonth(t *testing.T) {
	mt := DayOfMonth()
	_, err := time.Parse(DayOfMonthFormat, mt)
	if err != nil {
		t.Error("function DayOfMonth need return valid digit")
	}
}

func TestFakeTimestamp(t *testing.T) {

	tstmp := Timestamp()

	_, err := time.Parse(fmt.Sprintf("%s %s", BaseDateFormat, TimeFormat), tstmp)
	if err != nil {
		t.Error("function Timestamp need return valid timestamp format")
	}
}

func TestFakeCentury(t *testing.T) {
	centry := Century()
	if !slice.Contains(century, centry) {
		t.Error("Expected century from functuon Century")
	}
}

func TestFakeTimeZone(t *testing.T) {
	tz := Timezone()
	if !slice.Contains(timezones, tz) {
		t.Error("Expected timezone from variable timezones")
	}
}

func TestFakeTimePeriod(t *testing.T) {
	periode := Timeperiod()
	_, err := time.Parse(TimePeriodFormat, periode)
	if err != nil {
		t.Error("function TimePeriod need return valid period")
	}
}
