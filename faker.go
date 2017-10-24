package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits      = 6                    // 6 bits to represent a letter index
	letterIdxMask      = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax       = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	tagName            = "faker"
	Email              = "email"
	MacAddress         = "mac_address"
	DomainName         = "domain_name"
	UserName           = "username"
	Url                = "url"
	IPV4               = "ipv4"
	IPV6               = "ipv6"
	LATITUDE           = "lat"
	LONGITUDE          = "long"
	CREDIT_CARD_NUMBER = "cc_number"
	CREDIT_CARD_TYPE   = "cc_type"
)

var mapperTag = map[string]func() string{
	Email:              getNetworker().Email,
	MacAddress:         getNetworker().MacAddress,
	DomainName:         getNetworker().DomainName,
	Url:                getNetworker().Url,
	UserName:           getNetworker().UserName,
	IPV4:               getNetworker().Ipv4,
	IPV6:               getNetworker().Ipv6,
	CREDIT_CARD_TYPE:   getPayment().CreditCardType,
	CREDIT_CARD_NUMBER: getPayment().CreditCardNumber,
}

// Error when get fake from ptr
var ErrUnsupportedKindPtr = "Unsupported kind: %s Change Without using * (pointer) in Field of %s"

// Error when pass unsupported kind
var ErrUnsupportedKind = "Unsupported kind: %s"

// Error when value  is not pointer
var ErrValueNotPtr = "Not a pointer value"

// Error when tah not supported
var ErrTagNotSupported = "String Tag not unsupported"

// FakeData is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func FakeData(a interface{}) error {
	bootstrap()
	return setData(reflect.ValueOf(a))
}

// pre load
func bootstrap() {
	getNetworker()
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
		v.Set(reflect.ValueOf([]int{
			int(r.Intn(100)),
			int(r.Intn(100)),
			int(r.Intn(100)),
		}))
	case `[]int8`:
		v.Set(reflect.ValueOf([]int8{
			int8(r.Intn(100)),
			int8(r.Intn(100)),
			int8(r.Intn(100)),
		}))
	case `[]int16`:
		v.Set(reflect.ValueOf([]int16{
			int16(r.Intn(100)),
			int16(r.Intn(100)),
			int16(r.Intn(100)),
		}))
	case `[]int32`:
		v.Set(reflect.ValueOf([]int32{
			int32(r.Intn(100)),
			int32(r.Intn(100)),
			int32(r.Intn(100)),
		}))
	case `[]int64`:
		v.Set(reflect.ValueOf([]int64{
			int64(r.Intn(100)),
			int64(r.Intn(100)),
			int64(r.Intn(100)),
		}))
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

func setData(v reflect.Value) (err error) {
	r := rand.New(src)

	if v.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}

	v = reflect.Indirect(v)
	switch v.Kind() {

	case reflect.Int:
		v.Set(reflect.ValueOf(int(r.Intn(100))))
	case reflect.Int8:
		v.Set(reflect.ValueOf(int8(r.Intn(100))))
	case reflect.Int16:
		v.Set(reflect.ValueOf(int16(r.Intn(100))))
	case reflect.Int32:
		v.Set(reflect.ValueOf(int32(r.Intn(100))))
	case reflect.Int64:
		v.Set(reflect.ValueOf(int64(r.Intn(100))))
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

			t := v.Type()
			for i := 0; i < v.NumField(); i++ {
				tag := t.Field(i).Tag.Get(tagName)

				if tag != "" {
					err = setDataWithTag(v.Field(i).Addr(), tag)
				} else {
					err = setData(v.Field(i).Addr())
				}

				if err != nil {
					return err
				}
			}
		}
	case reflect.Ptr:
		return fmt.Errorf(ErrUnsupportedKindPtr, v.Kind().String(), v.Type().String())
	default:
		return fmt.Errorf(ErrUnsupportedKind, v.Kind().String())
	}

	return nil
}

func setDataWithTag(v reflect.Value, tag string) error {

	if v.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}

	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return userDefinedFloat(v, tag)
	case reflect.String:
		return userDefinedString(v, tag)
	case reflect.Slice:
		return setSliceData(v)
	}
	return nil
}

func userDefinedFloat(v reflect.Value, tag string) error {
	r := rand.New(src)
	kind := v.Kind()
	switch tag {
	case LATITUDE:
		val := r.Float32()*180 - 90
		if kind == reflect.Float32 {

			v.Set(reflect.ValueOf(val))
			return nil
		}
		v.Set(reflect.ValueOf(float64(val)))
		return nil
	case LONGITUDE:
		val := r.Float32()*360 - 180
		if kind == reflect.Float32 {
			v.Set(reflect.ValueOf(val))
			return nil
		}
		v.Set(reflect.ValueOf(float64(val)))
		return nil

	}
	return nil
}

func userDefinedString(v reflect.Value, tag string) error {
	val := ""
	if _, exist := mapperTag[tag]; !exist {
		return errors.New(ErrTagNotSupported)
	}
	val = mapperTag[tag]()
	v.SetString(val)
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

func randomElementFromSliceString(s []string) string {
	rand.Seed(time.Now().Unix())
	return s[rand.Int()%len(s)]
}
