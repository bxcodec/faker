package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var mu = &sync.Mutex{}

const (
	letterIdxBits      = 6                    // 6 bits to represent a letter index
	letterIdxMask      = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax       = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tagName            = "faker"
	Email              = "email"
	MacAddress         = "mac_address"
	DomainName         = "domain_name"
	UserName           = "username"
	Url                = "url"
	IPV4               = "ipv4"
	IPV6               = "ipv6"
	PASSWORD           = "password"
	LATITUDE           = "lat"
	LONGITUDE          = "long"
	CREDIT_CARD_NUMBER = "cc_number"
	CREDIT_CARD_TYPE   = "cc_type"
	PHONE_NUMBER       = "phone_number"
	TOLL_FREE_NUMBER   = "tool_free_number"
	E164_PHONE_NUMBER  = "e_164_phone_number"
	TITLE_MALE         = "title_male"
	TITLE_FEMALE       = "title_female"
	FIRST_NAME_MALE    = "first_name_male"
	FIRST_NAME_FEMALE  = "first_name_female"
	LAST_NAME          = "last_name"
	NAME               = "name"
	UNIX_TIME          = "unix_time"
	DATE               = "date"
	TIME               = "time"
	MONTH_NAME         = "month_name"
	YEAR               = "year"
	DAY_OF_WEEK        = "day_of_week"
	DAY_OF_MONTH       = "day_of_month"
	TIMESTAMP          = "timestamp"
	CENTURY            = "century"
	TIMEZONE           = "timezone"
	TIME_PERIOD        = "time_period"
	WORD               = "word"
	SENTENCE           = "sentence"
	SENTENCES          = "sentences"
)

var mapperTag = map[string]interface{}{
	Email:              GetNetworker().Email,
	MacAddress:         GetNetworker().MacAddress,
	DomainName:         GetNetworker().DomainName,
	Url:                GetNetworker().Url,
	UserName:           GetNetworker().UserName,
	IPV4:               GetNetworker().Ipv4,
	IPV6:               GetNetworker().Ipv6,
	PASSWORD:           GetNetworker().Password,
	CREDIT_CARD_TYPE:   GetPayment().CreditCardType,
	CREDIT_CARD_NUMBER: GetPayment().CreditCardNumber,
	LATITUDE:           GetAddress().Latitude,
	LONGITUDE:          GetAddress().Longitude,
	PHONE_NUMBER:       GetPhoner().PhoneNumber,
	TOLL_FREE_NUMBER:   GetPhoner().TollFreePhoneNumber,
	E164_PHONE_NUMBER:  GetPhoner().E164PhoneNumber,
	TITLE_MALE:         GetPerson().TitleMale,
	TITLE_FEMALE:       GetPerson().TitleFeMale,
	FIRST_NAME_MALE:    GetPerson().FirstNameMale,
	FIRST_NAME_FEMALE:  GetPerson().FirstNameFemale,
	LAST_NAME:          GetPerson().LastName,
	NAME:               GetPerson().Name,
	UNIX_TIME:          GetDateTimer().UnixTime,
	DATE:               GetDateTimer().Date,
	TIME:               GetDateTimer().Time,
	MONTH_NAME:         GetDateTimer().MonthName,
	YEAR:               GetDateTimer().Year,
	DAY_OF_WEEK:        GetDateTimer().DayOfWeek,
	DAY_OF_MONTH:       GetDateTimer().DayOfMonth,
	TIMESTAMP:          GetDateTimer().Timestamp,
	CENTURY:            GetDateTimer().Century,
	TIMEZONE:           GetDateTimer().TimeZone,
	TIME_PERIOD:        GetDateTimer().TimePeriod,
	WORD:               GetLorem().Word,
	SENTENCE:           GetLorem().Sentence,
	SENTENCES:          GetLorem().Sentences,
}

// Error when get fake from ptr
var ErrUnsupportedKindPtr = "Unsupported kind: %s Change Without using * (pointer) in Field of %s"

// Error when pass unsupported kind
var ErrUnsupportedKind = "Unsupported kind: %s"

// Error when value  is not pointer
var ErrValueNotPtr = "Not a pointer value"

// Error when tag not supported
var ErrTagNotSupported = "Tag unsupported"

// Error when passed more arguments
var ErrMoreArguments = "Passed more arguments than is possible : (%d)"

// FakeData is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func FakeData(a interface{}) error {

	reflectType := reflect.TypeOf(a)

	if reflectType.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}

	finalValue, err := getValue(reflectType.Elem())
	if err != nil {
		return err
	}
	rval := reflect.ValueOf(a)
	rval.Elem().Set(finalValue)
	return nil
}

func getValue(t reflect.Type) (reflect.Value, error) {
	r := rand.New(src)
	k := t.Kind()

	switch k {
	case reflect.Ptr:

		v := reflect.New(t.Elem())
		val, err := getValue(t.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		v.Elem().Set(val)
		return v, nil
	case reflect.Struct:

		switch t.String() {
		case "time.Time":
			ft := time.Unix(r.Int63(), 0)
			return reflect.ValueOf(ft), nil
		default:
			v := reflect.New(t).Elem()
			for i := 0; i < v.NumField(); i++ {
				tag := t.Field(i).Tag.Get(tagName)
				if tag == "" {
					val, err := getValue(v.Field(i).Type())
					if err != nil {
						return reflect.Value{}, err
					}
					v.Field(i).Set(val)
				} else {
					err := setDataWithTag(v.Field(i).Addr(), tag)
					if err != nil {
						return reflect.Value{}, err
					}
				}

			}
			return v, nil
		}

	case reflect.String:
		res := randomString(25)
		return reflect.ValueOf(res), nil
	case reflect.Array, reflect.Slice:
		len := r.Intn(100)
		v := reflect.MakeSlice(t, len, len)
		for i := 0; i < v.Len(); i++ {
			val, err := getValue(t.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Int:
		return reflect.ValueOf(int(r.Intn(100))), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(r.Intn(100))), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(r.Intn(100))), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(r.Intn(100))), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(r.Intn(100))), nil
	case reflect.Float32:
		return reflect.ValueOf(r.Float32()), nil
	case reflect.Float64:
		return reflect.ValueOf(r.Float64()), nil
	case reflect.Bool:
		val := r.Intn(2) > 0
		return reflect.ValueOf(val), nil

	case reflect.Uint:
		return reflect.ValueOf(uint(r.Intn(100))), nil

	case reflect.Uint8:
		return reflect.ValueOf(uint8(r.Intn(100))), nil

	case reflect.Uint16:
		return reflect.ValueOf(uint16(r.Intn(100))), nil

	case reflect.Uint32:
		return reflect.ValueOf(uint32(r.Intn(100))), nil

	case reflect.Uint64:
		return reflect.ValueOf(uint64(r.Intn(100))), nil

	case reflect.Map:
		v := reflect.MakeMap(t)
		len := r.Intn(100)
		for i := 0; i < len; i++ {
			key, err := getValue(t.Key())
			if err != nil {
				return reflect.Value{}, err
			}
			val, err := getValue(t.Elem())
			if err != nil {
				return reflect.Value{}, err
			}
			v.SetMapIndex(key, val)
		}
		return v, nil
	default:
		err := fmt.Errorf("no support for kind %+v", t)
		return reflect.Value{}, err
	}

}

// func setData(v reflect.Value) (err error) {
// 	r := rand.New(src)

// 	if v.Kind() != reflect.Ptr {
// 		return errors.New(ErrValueNotPtr)
// 	}

// 	v = reflect.Indirect(v)
// 	switch v.Kind() {

// 	case reflect.Int:
// 		v.Set(reflect.ValueOf(int(r.Intn(100))))
// 	case reflect.Int8:
// 		v.Set(reflect.ValueOf(int8(r.Intn(100))))
// 	case reflect.Int16:
// 		v.Set(reflect.ValueOf(int16(r.Intn(100))))
// 	case reflect.Int32:
// 		v.Set(reflect.ValueOf(int32(r.Intn(100))))
// 	case reflect.Int64:
// 		v.Set(reflect.ValueOf(int64(r.Intn(100))))
// 	case reflect.Float32:
// 		v.Set(reflect.ValueOf(r.Float32()))
// 	case reflect.Float64:
// 		v.Set(reflect.ValueOf(r.Float64()))
// 	case reflect.String:
// 		v.SetString(randomString(25))
// 	case reflect.Bool:
// 		val := r.Intn(2) > 0
// 		v.SetBool(val)
// 	case reflect.Slice:
// 		return setSliceData(v)
// 	case reflect.Struct:

// 		switch v.Type().String() {
// 		case "time.Time":

// 			ft := time.Unix(r.Int63(), 0)
// 			v.Set(reflect.ValueOf(ft))
// 		default:

// 			t := v.Type()
// 			for i := 0; i < v.NumField(); i++ {
// 				tag := t.Field(i).Tag.Get(tagName)

// 				if tag != "" {
// 					err = setDataWithTag(v.Field(i).Addr(), tag)
// 				} else {
// 					err = setData(v.Field(i).Addr())
// 				}

// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}

// 	case reflect.Ptr:
// 		// fmt.Println("Reflect Type", v.Type())

// 		// v = v.Addr().Elem().Elem()
// 		// // v = reflect.New(v.Type())
// 		// fmt.Println("Reflect Type", v.Type())
// 		// // err := setData(v)
// 		// if err != nil {
// 		// 	v.Set(reflect.Zero(v.Type()))
// 		// 	return nil
// 		// }
// 		// v.Set(reflect.New(v.Elem()))
// 		// v.Elem
// 		return fmt.Errorf(ErrUnsupportedKindPtr, v.Kind().String(), v.Type().String())
// 		// }

// 	default:
// 		return fmt.Errorf(ErrUnsupportedKind, v.Kind().String())
// 	}

// 	return nil
// }

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

	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16:
		return userDefinedInt(v, tag)
	}
	return nil
}

func userDefinedFloat(v reflect.Value, tag string) error {
	if _, exist := mapperTag[tag]; !exist {
		return errors.New(ErrTagNotSupported)
	}
	mapperTag[tag].(func(v reflect.Value) error)(v)
	return nil
}

func userDefinedString(v reflect.Value, tag string) error {
	val := ""
	if _, exist := mapperTag[tag]; !exist {
		return errors.New(ErrTagNotSupported)
	}
	val = mapperTag[tag].(func() string)()
	v.SetString(val)
	return nil
}

func userDefinedInt(v reflect.Value, tag string) error {
	if _, exist := mapperTag[tag]; !exist {
		return errors.New(ErrTagNotSupported)
	}
	mapperTag[tag].(func(v reflect.Value) error)(v)
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
func randomStringNumber(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(numberBytes) {
			b[i] = numberBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

/**
/ Get three parameters , only first mandatory and the rest are optional
/ --- If only set one parameter :  This means the minimum number of digits and the total number
/ --- If only set two parameters : First this is min digit and second max digit and the total number the difference between them
/ --- If only three parameters: the third argument set Max count Digit
*/
func RandomInt(parameters ...int) (p []int, err error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch len(parameters) {
	case 1:
		minCount := parameters[0]
		p = r.Perm(minCount)
		for i := range p {
			p[i] += minCount
		}
	case 2:
		minDigit, maxDigit := parameters[0], parameters[1]
		p = r.Perm(maxDigit - minDigit + 1)

		for i := range p {
			p[i] += minDigit
		}
	default:
		err = fmt.Errorf(ErrMoreArguments, len(parameters))
	}
	return p, err
}
