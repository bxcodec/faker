package faker

import (
	"math/rand"
	"reflect"
	"time"
)
const (
	BaseDate = "2006-01-02"
	Time = "15:04:05"
	Mounth = "January"
	Year = "2006"
)

type DateTimer interface {
	UnixTime(v reflect.Value) error
	Date() string
	Time() string
	Month() string
	Year() string
}

var date DateTimer

func getDateTimer() DateTimer {
	mu.Lock()
	defer mu.Unlock()

	if date == nil {
		date = &DateTime{}
	}
	return date
}

func SetDateTimer(d DateTimer) {
	date = d
}

type DateTime struct {
}

// get unix time
func (d DateTime) UnixTime(v reflect.Value) error {
	kind := v.Kind()

	if kind == reflect.Int64 {
		v.SetInt(RandomUnixTime())
	} else {
		v.SetInt(0)
	}
	return nil
}

// format example BaseDate const
func (d DateTime) Date() string {
	return time.Unix(RandomUnixTime(), 0).Format(BaseDate)
}

func (d DateTime) Time() string {
	return time.Unix(RandomUnixTime(), 0).Format(Time)
}

func (d DateTime) Month() string {
	return time.Unix(RandomUnixTime(), 0).Format(Mounth)
}

func (d DateTime) Year() string {
	return time.Unix(RandomUnixTime(), 0).Format(Year)
}


// helper function
func RandomUnixTime() int64 {
	return rand.Int63n(time.Now().Unix())
}
