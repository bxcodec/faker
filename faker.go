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
	Email:              getNetworker().Email,
	MacAddress:         getNetworker().MacAddress,
	DomainName:         getNetworker().DomainName,
	Url:                getNetworker().Url,
	UserName:           getNetworker().UserName,
	IPV4:               getNetworker().Ipv4,
	IPV6:               getNetworker().Ipv6,
	PASSWORD:           getNetworker().Password,
	CREDIT_CARD_TYPE:   getPayment().CreditCardType,
	CREDIT_CARD_NUMBER: getPayment().CreditCardNumber,
	LATITUDE:           getAddress().Latitude,
	LONGITUDE:          getAddress().Longitude,
	PHONE_NUMBER:       getPhoner().PhoneNumber,
	TOLL_FREE_NUMBER:   getPhoner().TollFreePhoneNumber,
	E164_PHONE_NUMBER:  getPhoner().E164PhoneNumber,
	TITLE_MALE:         getPerson().TitleMale,
	TITLE_FEMALE:       getPerson().TitleFeMale,
	FIRST_NAME_MALE:    getPerson().FirstNameMale,
	FIRST_NAME_FEMALE:  getPerson().FirstNameFemale,
	LAST_NAME:          getPerson().LastName,
	NAME:               getPerson().Name,
	UNIX_TIME:          getDateTimer().UnixTime,
	DATE:               getDateTimer().Date,
	TIME:               getDateTimer().Time,
	MONTH_NAME:         getDateTimer().MonthName,
	YEAR:               getDateTimer().Year,
	DAY_OF_WEEK:        getDateTimer().DayOfWeek,
	DAY_OF_MONTH:       getDateTimer().DayOfMonth,
	TIMESTAMP:          getDateTimer().Timestamp,
	CENTURY:            getDateTimer().Century,
	TIMEZONE:           getDateTimer().TimeZone,
	TIME_PERIOD:        getDateTimer().TimePeriod,
	WORD:               getLorem().Word,
	SENTENCE:           getLorem().Sentence,
	SENTENCES:          getLorem().Sentences,
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
