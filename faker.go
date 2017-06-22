package faker

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type SomeStruct struct {
	ID          int64
	Name        string
	Hobbies     []string
	Categories  []int64
	OtherStruct AStruct
}
type AStruct struct {
	Number int64
	Height int64
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func FakeData(a interface{}) error {

	return setData(reflect.ValueOf(a))
}

func setSliceData(v reflect.Value) error {
	v = reflect.Indirect(v)
	fmt.Println(v.Type())
	switch v.Type().String() {
	case `[]int64`:
		v.Set(reflect.ValueOf([]int64{1, 2, 3}))
		break
	case `[]string`:
		v.Set(reflect.ValueOf([]string{"hello", "world"}))
		break
	default:
		return errors.New("Slice of Struct Not Supported Yet")
	}
	return nil
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
		v.Set(reflect.ValueOf(int8(r.Int())))
	case reflect.Int16:
		v.Set(reflect.ValueOf(int16(r.Int())))
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
		val := r.Intn(1) > 0
		v.SetBool(val)
	case reflect.Slice:
		return setSliceData(v)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			err := setData(v.Field(i).Addr())
			if err != nil {
				return err
			}
		}

	default:
		return errors.New("Unsupported kind: " + v.Kind().String() + " Change Without using * (pointer) in any Field of Struct")

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
