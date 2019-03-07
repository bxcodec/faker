package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	mu = &sync.Mutex{}
	// Sets nil if the value type is struct or map and the size of it equals to zero.
	shouldSetNil = false
	//Sets random integer generation to zero for slice and maps
	testRandZero = false
	//Sets the default number of string when it is created randomly.
	randomStringLen = 25
	//Sets the boundary for random value generation. Boundaries can not exceed integer(4 byte...)
	nBoundary = numberBoundary{start: 0, end: 100}
	//Sets the random size for slices and maps.
	randomSize = 100
)

type numberBoundary struct {
	start int
	end   int
}

// Supported tags
const (
	letterIdxBits      = 6                    // 6 bits to represent a letter index
	letterIdxMask      = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax       = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tagName            = "faker"
	ID                 = "uuid_digit"
	HyphenatedID       = "uuid_hyphenated"
	EmailTag           = "email"
	MacAddressTag      = "mac_address"
	DomainNameTag      = "domain_name"
	UserNameTag        = "username"
	URLTag             = "url"
	IPV4Tag            = "ipv4"
	IPV6Tag            = "ipv6"
	PASSWORD           = "password"
	LATITUDE           = "lat"
	LONGITUDE          = "long"
	CreditCardNumber   = "cc_number"
	CreditCardType     = "cc_type"
	PhoneNumber        = "phone_number"
	TollFreeNumber     = "toll_free_number"
	E164PhoneNumber    = "e_164_phone_number"
	TitleMaleTag       = "title_male"
	TitleFemaleTag     = "title_female"
	FirstNameTag       = "first_name"
	FirstNameMaleTag   = "first_name_male"
	FirstNameFemaleTag = "first_name_female"
	LastNameTag        = "last_name"
	NAME               = "name"
	UnixTimeTag        = "unix_time"
	DATE               = "date"
	TIME               = "time"
	MonthNameTag       = "month_name"
	YEAR               = "year"
	DayOfWeekTag       = "day_of_week"
	DayOfMonthTag      = "day_of_month"
	TIMESTAMP          = "timestamp"
	CENTURY            = "century"
	TIMEZONE           = "timezone"
	TimePeriodTag      = "time_period"
	WORD               = "word"
	SENTENCE           = "sentence"
	PARAGRAPH          = "paragraph"
	Currency           = "currency"
	Amount             = "amount"
	AmountWithCurrency = "amount_with_currency"
	SKIP               = "-"
	Length             = "len"
	BoundaryStart      = "boundary_start"
	BoundaryEnd        = "boundary_end"
	Equals             = "="
	comma              = ","
)

var defaultTag = map[string]string{
	EmailTag:           EmailTag,
	MacAddressTag:      MacAddressTag,
	DomainNameTag:      DomainNameTag,
	URLTag:             URLTag,
	UserNameTag:        UserNameTag,
	IPV4Tag:            IPV4Tag,
	IPV6Tag:            IPV6Tag,
	PASSWORD:           PASSWORD,
	CreditCardType:     CreditCardType,
	CreditCardNumber:   CreditCardNumber,
	LATITUDE:           LATITUDE,
	LONGITUDE:          LONGITUDE,
	PhoneNumber:        PhoneNumber,
	TollFreeNumber:     TollFreeNumber,
	E164PhoneNumber:    E164PhoneNumber,
	TitleMaleTag:       TitleMaleTag,
	TitleFemaleTag:     TitleFemaleTag,
	FirstNameTag:       FirstNameTag,
	FirstNameMaleTag:   FirstNameMaleTag,
	FirstNameFemaleTag: FirstNameFemaleTag,
	LastNameTag:        LastNameTag,
	NAME:               NAME,
	UnixTimeTag:        UnixTimeTag,
	DATE:               DATE,
	TIME:               TimeFormat,
	MonthNameTag:       MonthNameTag,
	YEAR:               YearFormat,
	DayOfWeekTag:       DayOfWeekTag,
	DayOfMonthTag:      DayOfMonthFormat,
	TIMESTAMP:          TIMESTAMP,
	CENTURY:            CENTURY,
	TIMEZONE:           TIMEZONE,
	TimePeriodTag:      TimePeriodFormat,
	WORD:               WORD,
	SENTENCE:           SENTENCE,
	PARAGRAPH:          PARAGRAPH,
	Currency:           Currency,
	Amount:             Amount,
	AmountWithCurrency: AmountWithCurrency,
	ID:                 ID,
	HyphenatedID:       HyphenatedID,
}

// TaggedFunction used as the standard layout function for tag providers in struct.
// This type also can be used for custom provider.
type TaggedFunction func(v reflect.Value) (interface{}, error)

var mapperTag = map[string]TaggedFunction{
	EmailTag:           GetNetworker().Email,
	MacAddressTag:      GetNetworker().MacAddress,
	DomainNameTag:      GetNetworker().DomainName,
	URLTag:             GetNetworker().URL,
	UserNameTag:        GetNetworker().UserName,
	IPV4Tag:            GetNetworker().IPv4,
	IPV6Tag:            GetNetworker().IPv6,
	PASSWORD:           GetNetworker().Password,
	CreditCardType:     GetPayment().CreditCardType,
	CreditCardNumber:   GetPayment().CreditCardNumber,
	LATITUDE:           GetAddress().Latitude,
	LONGITUDE:          GetAddress().Longitude,
	PhoneNumber:        GetPhoner().PhoneNumber,
	TollFreeNumber:     GetPhoner().TollFreePhoneNumber,
	E164PhoneNumber:    GetPhoner().E164PhoneNumber,
	TitleMaleTag:       GetPerson().TitleMale,
	TitleFemaleTag:     GetPerson().TitleFeMale,
	FirstNameTag:       GetPerson().FirstName,
	FirstNameMaleTag:   GetPerson().FirstNameMale,
	FirstNameFemaleTag: GetPerson().FirstNameFemale,
	LastNameTag:        GetPerson().LastName,
	NAME:               GetPerson().Name,
	UnixTimeTag:        GetDateTimer().UnixTime,
	DATE:               GetDateTimer().Date,
	TIME:               GetDateTimer().Time,
	MonthNameTag:       GetDateTimer().MonthName,
	YEAR:               GetDateTimer().Year,
	DayOfWeekTag:       GetDateTimer().DayOfWeek,
	DayOfMonthTag:      GetDateTimer().DayOfMonth,
	TIMESTAMP:          GetDateTimer().Timestamp,
	CENTURY:            GetDateTimer().Century,
	TIMEZONE:           GetDateTimer().TimeZone,
	TimePeriodTag:      GetDateTimer().TimePeriod,
	WORD:               GetLorem().Word,
	SENTENCE:           GetLorem().Sentence,
	PARAGRAPH:          GetLorem().Paragraph,
	Currency:           GetPrice().Currency,
	Amount:             GetPrice().Amount,
	AmountWithCurrency: GetPrice().AmountWithCurrency,
	ID:                 GetIdentifier().Digit,
	HyphenatedID:       GetIdentifier().Hyphenated,
}

// Generic Error Messages for tags
// 		ErrUnsupportedKindPtr: Error when get fake from ptr
// 		ErrUnsupportedKind: Error on passing unsupported kind
// 		ErrValueNotPtr: Error when value is not pointer
// 		ErrTagNotSupported: Error when tag is not supported
// 		ErrTagAlreadyExists: Error when tag exists and call AddProvider
// 		ErrMoreArguments: Error on passing more arguments
// 		ErrNotSupportedPointer: Error when passing unsupported pointer
var (
	ErrUnsupportedKindPtr  = "Unsupported kind: %s Change Without using * (pointer) in Field of %s"
	ErrUnsupportedKind     = "Unsupported kind: %s"
	ErrValueNotPtr         = "Not a pointer value"
	ErrTagNotSupported     = "Tag unsupported"
	ErrTagAlreadyExists    = "Tag exists"
	ErrMoreArguments       = "Passed more arguments than is possible : (%d)"
	ErrNotSupportedPointer = "Use sample:=new(%s)\n faker.FakeData(sample) instead"
	ErrSmallerThanZero     = "Size:%d is smaller than zero."

	ErrStartValueBiggerThanEnd = "Start value can not be bigger than end value."
	ErrWrongFormattedTag       = "Tag \"%s\" is not written properly"
	ErrUnknownType             = "Unknown Type"
	ErrNotSupportedTypeForTag  = "Type is not supported by tag."
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SetNilIfLenIsZero allows to set nil for the slice and maps, if size is 0.
func SetNilIfLenIsZero(setNil bool) {
	shouldSetNil = setNil
}

// SetRandomStringLength sets a length for random string generation
func SetRandomStringLength(size int) error {
	if size < 0 {
		return fmt.Errorf(ErrSmallerThanZero, size)
	}
	randomStringLen = size
	return nil
}

// SetRandomMapAndSliceSize sets the size for maps and slices for random generation.
func SetRandomMapAndSliceSize(size int) error {
	if size < 0 {
		return fmt.Errorf(ErrSmallerThanZero, size)
	}
	randomSize = size
	return nil
}

// SetRandomNumberBoundaries sets boundary for random number generation
func SetRandomNumberBoundaries(start, end int) error {
	if start > end {
		return errors.New(ErrStartValueBiggerThanEnd)
	}
	nBoundary = numberBoundary{start: start, end: end}
	return nil
}

// FakeData is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func FakeData(a interface{}) error {

	reflectType := reflect.TypeOf(a)

	if reflectType.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}

	if reflect.ValueOf(a).IsNil() {
		return fmt.Errorf(ErrNotSupportedPointer, reflectType.Elem().String())
	}

	rval := reflect.ValueOf(a)

	finalValue, err := getValue(reflectType.Elem())
	if err != nil {
		return err
	}

	rval.Elem().Set(finalValue.Convert(reflectType.Elem()))
	return nil
}

// AddProvider extend faker with tag to generate fake data with specified custom algoritm
// Example:
// 		type Gondoruwo struct {
// 			Name       string
// 			Locatadata int
// 		}
//
// 		type Sample struct {
// 			ID                 int64     `faker:"customIdFaker"`
// 			Gondoruwo          Gondoruwo `faker:"gondoruwo"`
// 			Danger             string    `faker:"danger"`
// 		}
//
// 		func CustomGenerator() {
// 			// explicit
// 			faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
// 			 	return int64(43), nil
// 			})
// 			// functional
// 			faker.AddProvider("danger", func() faker.TaggedFunction {
// 				return func(v reflect.Value) (interface{}, error) {
// 					return "danger-ranger", nil
// 				}
// 			}())
// 			faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
// 				obj := Gondoruwo{
// 					Name:       "Power",
// 					Locatadata: 324,
// 				}
// 				return obj, nil
// 			})
// 		}
//
// 		func main() {
// 			CustomGenerator()
// 			var sample Sample
// 			faker.FakeData(&sample)
// 			fmt.Printf("%+v", sample)
// 		}
//
// Will print
// 		{ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger}
// Notes: when using a custom provider make sure to return the same type as the field
func AddProvider(tag string, provider TaggedFunction) error {
	if _, ok := mapperTag[tag]; ok {
		return errors.New(ErrTagAlreadyExists)
	}

	mapperTag[tag] = provider

	return nil
}

func getValue(t reflect.Type) (reflect.Value, error) {
	k := t.Kind()

	switch k {
	case reflect.Ptr:
		v := reflect.New(t.Elem())
		val, err := getValue(t.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		v.Elem().Set(val.Convert(t.Elem()))
		return v, nil
	case reflect.Struct:

		switch t.String() {
		case "time.Time":
			ft := time.Now().Add(time.Duration(rand.Int63()))
			return reflect.ValueOf(ft), nil
		default:
			v := reflect.New(t).Elem()
			for i := 0; i < v.NumField(); i++ {
				if !v.Field(i).CanSet() {
					continue // to avoid panic to set on unexported field in struct
				}
				tag := t.Field(i).Tag.Get(tagName)

				switch tag {
				case "":
					val, err := getValue(v.Field(i).Type())
					if err != nil {
						return reflect.Value{}, err
					}
					val = val.Convert(v.Field(i).Type())
					v.Field(i).Set(val)
				case SKIP:
					continue
				default:
					err := setDataWithTag(v.Field(i).Addr(), tag)
					if err != nil {
						return reflect.Value{}, err
					}
				}

			}
			return v, nil
		}

	case reflect.String:
		res := randomString(randomStringLen)
		return reflect.ValueOf(res), nil
	case reflect.Array, reflect.Slice:
		len := randomSliceAndMapSize()
		if shouldSetNil && len == 0 {
			return reflect.Zero(t), nil
		}
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
		return reflect.ValueOf(randomInteger()), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(randomInteger())), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(randomInteger())), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(randomInteger())), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(randomInteger())), nil
	case reflect.Float32:
		return reflect.ValueOf(rand.Float32()), nil
	case reflect.Float64:
		return reflect.ValueOf(rand.Float64()), nil
	case reflect.Bool:
		val := rand.Intn(2) > 0
		return reflect.ValueOf(val), nil

	case reflect.Uint:
		return reflect.ValueOf(uint(randomInteger())), nil

	case reflect.Uint8:
		return reflect.ValueOf(uint8(randomInteger())), nil

	case reflect.Uint16:
		return reflect.ValueOf(uint16(randomInteger())), nil

	case reflect.Uint32:
		return reflect.ValueOf(uint32(randomInteger())), nil

	case reflect.Uint64:
		return reflect.ValueOf(uint64(randomInteger())), nil

	case reflect.Map:
		len := randomSliceAndMapSize()
		if shouldSetNil && len == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeMap(t)
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

func setDataWithTag(v reflect.Value, tag string) error {

	if v.Kind() != reflect.Ptr {
		return errors.New(ErrValueNotPtr)
	}
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.Ptr:
		if _, exist := mapperTag[tag]; !exist {
			return errors.New(ErrTagNotSupported)
		}
		if _, def := defaultTag[tag]; !def {
			res, err := mapperTag[tag](v)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(res))
			return nil
		}

		t := v.Type()
		newv := reflect.New(t.Elem())
		res, err := mapperTag[tag](newv.Elem())
		if err != nil {
			return err
		}
		rval := reflect.ValueOf(res)
		newv.Elem().Set(rval)
		v.Set(newv)
		return nil
	case reflect.String:
		return userDefinedString(v, tag)
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return userDefinedNumber(v, tag)
	case reflect.Slice, reflect.Array:
		return userDefinedArray(v, tag)
	case reflect.Map:
		return userDefinedMap(v, tag)
	default:
		if _, exist := mapperTag[tag]; !exist {
			return errors.New(ErrTagNotSupported)
		}
		res, err := mapperTag[tag](v)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(res))
	}
	return nil
}

func userDefinedMap(v reflect.Value, tag string) error {
	len := randomSliceAndMapSize()
	if shouldSetNil && len == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	definedMap := reflect.MakeMap(v.Type())
	for i := 0; i < len; i++ {
		key, err := getValueWithTag(v.Type().Key(), tag)
		if err != nil {
			return err
		}
		val, err := getValueWithTag(v.Type().Elem(), tag)
		if err != nil {
			return err
		}
		definedMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	v.Set(definedMap)
	return nil
}

func getValueWithTag(t reflect.Type, tag string) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res, err := extractNumberFromTag(tag, t)
		if err != nil {
			return nil, err
		}
		return res, nil
	case reflect.String:
		res, err := extractStringFromTag(tag)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return 0, errors.New(ErrUnknownType)
	}
}

func userDefinedArray(v reflect.Value, tag string) error {
	len := randomSliceAndMapSize()
	if shouldSetNil && len == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	array := reflect.MakeSlice(v.Type(), len, len)
	for i := 0; i < len; i++ {
		res, err := getValueWithTag(v.Type().Elem(), tag)
		if err != nil {
			return err
		}
		array.Index(i).Set(reflect.ValueOf(res))
	}
	v.Set(array)
	return nil
}

func userDefinedString(v reflect.Value, tag string) error {
	var res interface{}
	var err error

	if tagFunc, ok := mapperTag[tag]; ok {
		res, err = tagFunc(v)
		if err != nil {
			return err
		}
	} else {
		res, err = extractStringFromTag(tag)
		if err != nil {
			return err
		}
	}
	if res == nil {
		return errors.New(ErrTagNotSupported)
	}
	val, _ := res.(string)
	v.SetString(val)
	return nil
}

func userDefinedNumber(v reflect.Value, tag string) error {
	var res interface{}
	var err error

	if tagFunc, ok := mapperTag[tag]; ok {
		res, err = tagFunc(v)
		if err != nil {
			return err
		}
	} else {
		res, err = extractNumberFromTag(tag, v.Type())
		if err != nil {
			return err
		}
	}
	if res == nil {
		return errors.New(ErrTagNotSupported)
	}

	v.Set(reflect.ValueOf(res))
	return nil
}

func extractStringFromTag(tag string) (interface{}, error) {
	if !strings.Contains(tag, Length) {
		return nil, errors.New(ErrTagNotSupported)
	}
	len, err := extractNumberFromText(tag)
	if err != nil {
		return nil, err
	}
	res := randomString(len)
	return res, nil
}

func extractNumberFromTag(tag string, t reflect.Type) (interface{}, error) {
	if !strings.Contains(tag, BoundaryStart) || !strings.Contains(tag, BoundaryEnd) {
		return nil, errors.New(ErrTagNotSupported)
	}
	valuesStr := strings.SplitN(tag, comma, -1)
	if len(valuesStr) != 2 {
		return nil, fmt.Errorf(ErrWrongFormattedTag, tag)
	}
	startBoundary, err := extractNumberFromText(valuesStr[0])
	if err != nil {
		return nil, err
	}
	endBoundary, err := extractNumberFromText(valuesStr[1])
	if err != nil {
		return nil, err
	}
	boundary := numberBoundary{start: startBoundary, end: endBoundary}
	switch t.Kind() {
	case reflect.Uint:
		return uint(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint8:
		return uint8(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint16:
		return uint16(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint32:
		return uint32(randomIntegerWithBoundary(boundary)), nil
	case reflect.Uint64:
		return uint64(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int:
		return randomIntegerWithBoundary(boundary), nil
	case reflect.Int8:
		return int8(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int16:
		return int16(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int32:
		return int32(randomIntegerWithBoundary(boundary)), nil
	case reflect.Int64:
		return int64(randomIntegerWithBoundary(boundary)), nil
	default:
		return nil, errors.New(ErrNotSupportedTypeForTag)
	}
}

func extractNumberFromText(text string) (int, error) {
	text = strings.TrimSpace(text)
	texts := strings.SplitN(text, Equals, -1)
	if len(texts) != 2 {
		return 0, fmt.Errorf(ErrWrongFormattedTag, text)
	}
	return strconv.Atoi(texts[1])
}

func randomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
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

// randomIntegerWithBoundary returns a random integer between input start and end boundary. [start, end)
func randomIntegerWithBoundary(boundary numberBoundary) int {
	return rand.Intn(boundary.end-boundary.start) + boundary.start
}

// randomInteger returns a random integer between start and end boundary. [start, end)
func randomInteger() int {
	return rand.Intn(nBoundary.end-nBoundary.start) + nBoundary.start
}

// randomSliceAndMapSize returns a random integer between [0,randomSliceAndMapSize). If the testRandZero is set, returns 0
// Written for test purposes for shouldSetNil
func randomSliceAndMapSize() int {
	if testRandZero {
		return 0
	}
	return rand.Intn(randomSize)
}

func randomElementFromSliceString(s []string) string {
	return s[rand.Int()%len(s)]
}
func randomStringNumber(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
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

// RandomInt Get three parameters , only first mandatory and the rest are optional
// 		If only set one parameter :  This means the minimum number of digits and the total number
// 		If only set two parameters : First this is min digit and second max digit and the total number the difference between them
// 		If only three parameters: the third argument set Max count Digit
func RandomInt(parameters ...int) (p []int, err error) {
	switch len(parameters) {
	case 1:
		minCount := parameters[0]
		p = rand.Perm(minCount)
		for i := range p {
			p[i] += minCount
		}
	case 2:
		minDigit, maxDigit := parameters[0], parameters[1]
		p = rand.Perm(maxDigit - minDigit + 1)

		for i := range p {
			p[i] += minDigit
		}
	default:
		err = fmt.Errorf(ErrMoreArguments, len(parameters))
	}
	return p, err
}
