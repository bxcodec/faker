package faker

import (
	"math/rand"
	"reflect"
	"time"
)

type DateTimer interface {
	UnixTime(v reflect.Value) error
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

func (d DateTime) UnixTime(v reflect.Value) error {
	kind := v.Kind()

	if kind == reflect.Int64 {
		v.SetInt(rand.Int63n(time.Now().Unix()))
	} else {
		v.SetInt(0)
	}
	return nil
}
