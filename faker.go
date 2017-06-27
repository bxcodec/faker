package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// FakeData is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func FakeData(a interface{}) error {

	return setData(reflect.ValueOf(a))
}

func setSliceData(v reflect.Value) error {
	r := rand.New(src)
	v = reflect.Indirect(v)

	var err error

	switch v.Type().String() {
	case `[]bool`:
		val1 := r.Intn(2) > 0
		val2 := r.Intn(2) > 0
		val3 := r.Intn(2) > 0
		v.Set(reflect.ValueOf([]bool{val1, val2, val3}))
	case `[]int`:
		v.Set(reflect.ValueOf([]int{r.Int(), r.Int(), r.Int()}))
	case `[]int8`:
		v.Set(reflect.ValueOf([]int8{int8(r.Int()), int8(r.Int()), int8(r.Int())}))
	case `[]int16`:
		v.Set(reflect.ValueOf([]int16{int16(r.Int()), int16(r.Int()), int16(r.Int())}))
	case `[]int32`:
		v.Set(reflect.ValueOf([]int32{r.Int31(), r.Int31(), r.Int31()}))
	case `[]int64`:
		v.Set(reflect.ValueOf([]int64{r.Int63(), r.Int63(), r.Int63()}))
	case `[]float32`:
		v.Set(reflect.ValueOf([]float32{r.Float32(), r.Float32(), r.Float32()}))
	case `[]float64`:
		v.Set(reflect.ValueOf([]float64{r.Float64(), r.Float64(), r.Float64()}))
	case `[]string`:
		v.Set(reflect.ValueOf([]string{randomString(5), randomString(7)}))
	case `[]time.Time`:
		ft := time.Unix(r.Int63(), 0)
		ft2 := time.Unix(r.Int63(), 0)
		v.Set(reflect.ValueOf([]time.Time{ft, ft2}))

	default:
		err = errors.New("Slice of " + v.Type().String() + " Not Supported Yet")
	}
	return err
}
func setData(v reflect.Value) error {
	r := rand.New(src)

	if v.Kind() != reflect.Ptr {
		return errors.New("Not a pointer value")
	}

	v = reflect.Indirect(v)
	switch v.Kind() {

	case reflect.Int:
		v.SetInt(r.Int63())
	case reflect.Int8:
		v.Set(reflect.ValueOf(int8(r.Intn(8))))
	case reflect.Int16:
		v.Set(reflect.ValueOf(int16(r.Intn(16))))
	case reflect.Int32:
		v.Set(reflect.ValueOf(r.Int31()))
	case reflect.Int64:
		v.Set(reflect.ValueOf(r.Int63()))
	case reflect.Float32:
		v.Set(reflect.ValueOf(r.Float32()))
	case reflect.Float64:
		v.Set(reflect.ValueOf(r.Float64()))
	case reflect.String:
		v.SetString(randomString(25))
	case reflect.Bool:
		val := r.Intn(2) > 0
		v.SetBool(val)
	case reflect.Slice:
		return setSliceData(v)
	case reflect.Struct:

		if v.Type().String() == "time.Time" {
			ft := time.Unix(r.Int63(), 0)
			v.Set(reflect.ValueOf(ft))
		} else {
			for i := 0; i < v.NumField(); i++ {
				err := setData(v.Field(i).Addr())
				if err != nil {
					return err
				}
			}
		}

	case reflect.Ptr:
		return errors.New("Unsupported kind: " + v.Kind().String() + " Change Without using * (pointer) in Field of " + v.Type().String())
	default:
		return errors.New("Unsupported kind: " + v.Kind().String())
	}

	return nil
}

func randomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
