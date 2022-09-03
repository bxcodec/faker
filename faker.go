package faker

// Faker is a simple fake data generator for your own struct.
// Save your time, and Fake your data for your testing now.
import (
	cryptorand "crypto/rand"
	"errors"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	fakerErrors "github.com/bxcodec/faker/v4/pkg/errors"
	"github.com/bxcodec/faker/v4/pkg/interfaces"
	"github.com/bxcodec/faker/v4/pkg/options"
	"github.com/bxcodec/faker/v4/pkg/slice"
)

var (
	// Unique values are kept in memory so the generator retries if the value already exists
	uniqueValues = map[string][]interface{}{}
)

// Supported tags
const (
	letterIdxBits         = 6                    // 6 bits to represent a letter index
	letterIdxMask         = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax          = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	maxRetry              = 10000                // max number of retry for unique values
	tagName               = "faker"
	keep                  = "keep"
	unique                = "unique"
	ID                    = "uuid_digit"
	HyphenatedID          = "uuid_hyphenated"
	EmailTag              = "email"
	MacAddressTag         = "mac_address"
	DomainNameTag         = "domain_name"
	UserNameTag           = "username"
	URLTag                = "url"
	IPV4Tag               = "ipv4"
	IPV6Tag               = "ipv6"
	PASSWORD              = "password"
	JWT                   = "jwt"
	LATITUDE              = "lat"
	LONGITUDE             = "long"
	CreditCardNumber      = "cc_number"
	CreditCardType        = "cc_type"
	PhoneNumber           = "phone_number"
	TollFreeNumber        = "toll_free_number"
	E164PhoneNumberTag    = "e_164_phone_number"
	TitleMaleTag          = "title_male"
	TitleFemaleTag        = "title_female"
	FirstNameTag          = "first_name"
	FirstNameMaleTag      = "first_name_male"
	FirstNameFemaleTag    = "first_name_female"
	LastNameTag           = "last_name"
	NAME                  = "name"
	ChineseFirstNameTag   = "chinese_first_name"
	ChineseLastNameTag    = "chinese_last_name"
	ChineseNameTag        = "chinese_name"
	GENDER                = "gender"
	UnixTimeTag           = "unix_time"
	DATE                  = "date"
	TIME                  = "time"
	MonthNameTag          = "month_name"
	YEAR                  = "year"
	DayOfWeekTag          = "day_of_week"
	DayOfMonthTag         = "day_of_month"
	TIMESTAMP             = "timestamp"
	CENTURY               = "century"
	TIMEZONE              = "timezone"
	TimePeriodTag         = "time_period"
	WORD                  = "word"
	SENTENCE              = "sentence"
	PARAGRAPH             = "paragraph"
	CurrencyTag           = "currency"
	AmountTag             = "amount"
	AmountWithCurrencyTag = "amount_with_currency"
	SKIP                  = "-"
	Length                = "len"
	SliceLength           = "slice_len"
	Language              = "lang"
	BoundaryStart         = "boundary_start"
	BoundaryEnd           = "boundary_end"
	Equals                = "="
	comma                 = ","
	colon                 = ":"
	ONEOF                 = "oneof"
)

// PriorityTags define the priority order of the tag
var PriorityTags = []string{ID, HyphenatedID, EmailTag, MacAddressTag, DomainNameTag, UserNameTag, URLTag, IPV4Tag,
	IPV6Tag, PASSWORD, JWT, LATITUDE, LONGITUDE, CreditCardNumber, CreditCardType, PhoneNumber, TollFreeNumber,
	E164PhoneNumberTag, TitleMaleTag, TitleFemaleTag, FirstNameTag, FirstNameMaleTag, FirstNameFemaleTag, LastNameTag,
	NAME, ChineseFirstNameTag, ChineseLastNameTag, ChineseNameTag, GENDER, UnixTimeTag, DATE, TIME, MonthNameTag,
	YEAR, DayOfWeekTag, DayOfMonthTag, TIMESTAMP, CENTURY, TIMEZONE, TimePeriodTag, WORD, SENTENCE, PARAGRAPH,
	CurrencyTag, AmountTag, AmountWithCurrencyTag, SKIP, Length, SliceLength, Language, BoundaryStart, BoundaryEnd, ONEOF,
}

var defaultTag = map[string]string{
	EmailTag:              EmailTag,
	MacAddressTag:         MacAddressTag,
	DomainNameTag:         DomainNameTag,
	URLTag:                URLTag,
	UserNameTag:           UserNameTag,
	IPV4Tag:               IPV4Tag,
	IPV6Tag:               IPV6Tag,
	PASSWORD:              PASSWORD,
	JWT:                   JWT,
	CreditCardType:        CreditCardType,
	CreditCardNumber:      CreditCardNumber,
	LATITUDE:              LATITUDE,
	LONGITUDE:             LONGITUDE,
	PhoneNumber:           PhoneNumber,
	TollFreeNumber:        TollFreeNumber,
	E164PhoneNumberTag:    E164PhoneNumberTag,
	TitleMaleTag:          TitleMaleTag,
	TitleFemaleTag:        TitleFemaleTag,
	FirstNameTag:          FirstNameTag,
	FirstNameMaleTag:      FirstNameMaleTag,
	FirstNameFemaleTag:    FirstNameFemaleTag,
	LastNameTag:           LastNameTag,
	NAME:                  NAME,
	ChineseFirstNameTag:   ChineseFirstNameTag,
	ChineseLastNameTag:    ChineseLastNameTag,
	ChineseNameTag:        ChineseNameTag,
	GENDER:                GENDER,
	UnixTimeTag:           UnixTimeTag,
	DATE:                  DATE,
	TIME:                  TimeFormat,
	MonthNameTag:          MonthNameTag,
	YEAR:                  YearFormat,
	DayOfWeekTag:          DayOfWeekTag,
	DayOfMonthTag:         DayOfMonthFormat,
	TIMESTAMP:             TIMESTAMP,
	CENTURY:               CENTURY,
	TIMEZONE:              TIMEZONE,
	TimePeriodTag:         TimePeriodFormat,
	WORD:                  WORD,
	SENTENCE:              SENTENCE,
	PARAGRAPH:             PARAGRAPH,
	CurrencyTag:           CurrencyTag,
	AmountTag:             AmountTag,
	AmountWithCurrencyTag: AmountWithCurrencyTag,
	ID:                    ID,
	HyphenatedID:          HyphenatedID,
}

var mapperTag = map[string]interfaces.TaggedFunction{
	CreditCardType:        GetPayment().CreditCardType,
	CreditCardNumber:      GetPayment().CreditCardNumber,
	LATITUDE:              GetAddress().Latitude,
	LONGITUDE:             GetAddress().Longitude,
	PhoneNumber:           GetPhoner().PhoneNumber,
	TollFreeNumber:        GetPhoner().TollFreePhoneNumber,
	E164PhoneNumberTag:    GetPhoner().E164PhoneNumber,
	TitleMaleTag:          GetPerson().TitleMale,
	TitleFemaleTag:        GetPerson().TitleFeMale,
	FirstNameTag:          GetPerson().FirstName,
	FirstNameMaleTag:      GetPerson().FirstNameMale,
	FirstNameFemaleTag:    GetPerson().FirstNameFemale,
	LastNameTag:           GetPerson().LastName,
	NAME:                  GetPerson().Name,
	ChineseFirstNameTag:   GetPerson().ChineseFirstName,
	ChineseLastNameTag:    GetPerson().ChineseLastName,
	ChineseNameTag:        GetPerson().ChineseName,
	GENDER:                GetPerson().Gender,
	UnixTimeTag:           GetDateTimer().UnixTime,
	DATE:                  GetDateTimer().Date,
	TIME:                  GetDateTimer().Time,
	MonthNameTag:          GetDateTimer().MonthName,
	YEAR:                  GetDateTimer().Year,
	DayOfWeekTag:          GetDateTimer().DayOfWeek,
	DayOfMonthTag:         GetDateTimer().DayOfMonth,
	TIMESTAMP:             GetDateTimer().Timestamp,
	CENTURY:               GetDateTimer().Century,
	TIMEZONE:              GetDateTimer().TimeZone,
	TimePeriodTag:         GetDateTimer().TimePeriod,
	WORD:                  GetLorem().Word,
	SENTENCE:              GetLorem().Sentence,
	PARAGRAPH:             GetLorem().Paragraph,
	CurrencyTag:           GetPrice().Currency,
	AmountTag:             GetPrice().Amount,
	AmountWithCurrencyTag: GetPrice().AmountWithCurrency,
	ID:                    GetIdentifier().Digit,
	HyphenatedID:          GetIdentifier().Hyphenated,
}

// Compiled regexp
var (
	findLangReg     *regexp.Regexp
	findLenReg      *regexp.Regexp
	findSliceLenReg *regexp.Regexp
)

func init() {
	rand = mathrand.New(NewSafeSource(mathrand.NewSource(time.Now().UnixNano())))
	crypto = cryptorand.Reader
}

func init() {
	findLangReg, _ = regexp.Compile("lang=[a-z]{3}")
	findLenReg, _ = regexp.Compile(`len=\d+`)
	findSliceLenReg, _ = regexp.Compile(`slice_len=\d+`)

	randNameFlag = rand.Intn(100) // for person
}

// ResetUnique is used to forget generated unique values.
// Call this when you're done generating a dataset.
func ResetUnique() {
	uniqueValues = map[string][]interface{}{}
}

var (
	SetGenerateUniqueValues     = options.SetGenerateUniqueValues
	SetIgnoreInterface          = options.SetIgnoreInterface
	SetRandomStringLength       = options.SetRandomStringLength
	SetStringLang               = options.SetStringLang
	SetRandomMapAndSliceSize    = options.SetRandomMapAndSliceSize
	SetRandomMapAndSliceMaxSize = options.SetRandomMapAndSliceMaxSize
	SetRandomMapAndSliceMinSize = options.SetRandomMapAndSliceMinSize
	SetRandomNumberBoundaries   = options.SetRandomNumberBoundaries
)

func initMapperTagWithOption(opts ...options.OptionFunc) {
	mapperTag[EmailTag] = GetNetworker(opts...).Email
	mapperTag[MacAddressTag] = GetNetworker(opts...).MacAddress
	mapperTag[DomainNameTag] = GetNetworker(opts...).DomainName
	mapperTag[URLTag] = GetNetworker(opts...).URL
	mapperTag[UserNameTag] = GetNetworker(opts...).UserName
	mapperTag[IPV4Tag] = GetNetworker(opts...).IPv4
	mapperTag[IPV6Tag] = GetNetworker(opts...).IPv6
	mapperTag[PASSWORD] = GetNetworker(opts...).Password
	mapperTag[JWT] = GetNetworker(opts...).Jwt
}

func initOption(opt ...options.OptionFunc) *options.Options {
	opts := options.BuildOptions(opt)
	initMapperTagWithOption(opt...)
	return opts
}

// FakeData is the main function. Will generate a fake data based on your struct.  You can use this for automation testing, or anything that need automated data.
// You don't need to Create your own data for your testing.
func FakeData(a interface{}, opt ...options.OptionFunc) error {
	opts := initOption(opt...)
	reflectType := reflect.TypeOf(a)

	if reflectType.Kind() != reflect.Ptr {
		return errors.New(fakerErrors.ErrValueNotPtr)
	}

	if reflect.ValueOf(a).IsNil() {
		return fmt.Errorf(fakerErrors.ErrNotSupportedPointer, reflectType.Elem().String())
	}

	rval := reflect.ValueOf(a)
	finalValue, err := getFakedValue(a, opts)
	if err != nil {
		return err
	}

	rval.Elem().Set(finalValue.Elem().Convert(reflectType.Elem()))
	return nil
}

// AddProvider extend faker with tag to generate fake data with specified custom algorithm
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
func AddProvider(tag string, provider interfaces.TaggedFunction) error {
	if _, ok := mapperTag[tag]; ok {
		return errors.New(fakerErrors.ErrTagAlreadyExists)
	}
	PriorityTags = append(PriorityTags, tag)
	mapperTag[tag] = provider

	return nil
}

// RemoveProvider removes existing customization added with AddProvider
func RemoveProvider(tag string) error {
	if _, ok := mapperTag[tag]; !ok {
		return errors.New(fakerErrors.ErrTagDoesNotExist)
	}

	delete(mapperTag, tag)

	return nil
}

func getFakedValue(a interface{}, opts *options.Options) (reflect.Value, error) {
	t := reflect.TypeOf(a)
	if t == nil {
		if opts.IgnoreInterface {
			return reflect.New(reflect.TypeOf(reflect.Struct)), nil
		}
		return reflect.Value{}, fmt.Errorf("interface{} not allowed")
	}
	if opts.MaxDepthOption.RecursionOutOfLimit(t) {
		return reflect.Zero(t), nil
	}
	opts.MaxDepthOption.RememberType(t)
	defer func() {
		opts.MaxDepthOption.ForgetType(t)
	}()
	k := t.Kind()

	switch k {
	case reflect.Ptr:
		v := reflect.New(t.Elem())
		var val reflect.Value
		var err error
		if a != reflect.Zero(reflect.TypeOf(a)).Interface() {
			val, err = getFakedValue(reflect.ValueOf(a).Elem().Interface(), opts)
		} else {
			val, err = getFakedValue(v.Elem().Interface(), opts)
		}
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
			originalDataVal := reflect.ValueOf(a)
			v := reflect.New(t).Elem()
			retry := 0 // error if cannot generate unique value after maxRetry tries
			for i := 0; i < v.NumField(); i++ {
				if !v.Field(i).CanSet() {
					continue // to avoid panic to set on unexported field in struct
				}

				if _, ok := opts.IgnoreFields[t.Field(i).Name]; ok {
					continue
				}

				if p, ok := opts.FieldProviders[t.Field(i).Name]; ok {
					val, err := p()
					if err != nil {
						return reflect.Value{}, fmt.Errorf("custom provider for field %s: %w", t.Field(i).Name, err)
					}
					v.Field(i).Set(reflect.ValueOf(val))
					continue
				}

				tags := decodeTags(t, i)
				switch {
				case tags.keepOriginal:
					zero, err := isZero(reflect.ValueOf(a).Field(i))
					if err != nil {
						return reflect.Value{}, err
					}
					if zero {
						err := setDataWithTag(v.Field(i).Addr(), tags.fieldType, *opts)
						if err != nil {
							return reflect.Value{}, err
						}
						continue
					}
					v.Field(i).Set(reflect.ValueOf(a).Field(i))
				case tags.fieldType == "":
					val, err := getFakedValue(v.Field(i).Interface(), opts)
					if err != nil {
						return reflect.Value{}, err
					}
					val = val.Convert(v.Field(i).Type())
					v.Field(i).Set(val)
				case tags.fieldType == SKIP:
					item := originalDataVal.Field(i).Interface()
					if v.CanSet() && item != nil {
						v.Field(i).Set(reflect.ValueOf(item))
					}
				default:
					err := setDataWithTag(v.Field(i).Addr(), tags.fieldType, *opts)
					if err != nil {
						return reflect.Value{}, err
					}
				}

				if tags.unique {

					if retry >= maxRetry {
						return reflect.Value{}, fmt.Errorf(fakerErrors.ErrUniqueFailure, reflect.TypeOf(a).Field(i).Name)
					}

					value := v.Field(i).Interface()
					if slice.ContainsValue(uniqueValues[tags.fieldType], value) { // Retry if unique value already found
						i--
						retry++
						continue
					}
					retry = 0
					uniqueValues[tags.fieldType] = append(uniqueValues[tags.fieldType], value)
				} else {
					retry = 0
				}

			}
			return v, nil
		}

	case reflect.String:
		res, err := randomString(opts.RandomStringLength, *opts)
		return reflect.ValueOf(res), err
	case reflect.Slice:
		length := randomSliceAndMapSize(*opts)
		if opts.SetSliceMapNilIfLenZero && length == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeSlice(t, length, length)
		for i := 0; i < v.Len(); i++ {
			val, err := getFakedValue(v.Index(i).Interface(), opts)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Index(i).Type())
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Array:
		v := reflect.New(t).Elem()
		for i := 0; i < v.Len(); i++ {
			val, err := getFakedValue(v.Index(i).Interface(), opts)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Index(i).Type())
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Int:
		return reflect.ValueOf(randomInteger(opts)), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(randomInteger(opts))), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(randomInteger(opts))), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(randomInteger(opts))), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(randomInteger(opts))), nil
	case reflect.Float32:
		return reflect.ValueOf(float32(randomFloat(opts))), nil
	case reflect.Float64:
		return reflect.ValueOf(randomFloat(opts)), nil
	case reflect.Bool:
		val := rand.Intn(2) > 0
		return reflect.ValueOf(val), nil

	case reflect.Uint:
		return reflect.ValueOf(uint(randomInteger(opts))), nil

	case reflect.Uint8:
		return reflect.ValueOf(uint8(randomInteger(opts))), nil

	case reflect.Uint16:
		return reflect.ValueOf(uint16(randomInteger(opts))), nil

	case reflect.Uint32:
		return reflect.ValueOf(uint32(randomInteger(opts))), nil

	case reflect.Uint64:
		return reflect.ValueOf(uint64(randomInteger(opts))), nil

	case reflect.Map:
		length := randomSliceAndMapSize(*opts)
		if opts.SetSliceMapNilIfLenZero && length == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeMap(t)
		for i := 0; i < length; i++ {
			keyInstance := reflect.New(t.Key()).Elem().Interface()
			key, err := getFakedValue(keyInstance, opts)
			if err != nil {
				return reflect.Value{}, err
			}

			valueInstance := reflect.New(t.Elem()).Elem().Interface()
			val, err := getFakedValue(valueInstance, opts)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(v.Type().Elem())
			v.SetMapIndex(key, val)
		}
		return v, nil
	default:
		err := fmt.Errorf("no support for kind %+v", t)
		return reflect.Value{}, err
	}

}

func isZero(field reflect.Value) (bool, error) {
	if field.Kind() == reflect.Map {
		return field.Len() == 0, nil
	}

	for _, kind := range []reflect.Kind{reflect.Struct, reflect.Slice, reflect.Array} {
		if kind == field.Kind() {
			return false, fmt.Errorf("keep not allowed on struct")
		}
	}
	return reflect.Zero(field.Type()).Interface() == field.Interface(), nil
}

func decodeTags(typ reflect.Type, i int) structTag {
	tagField := typ.Field(i).Tag.Get(tagName)
	tags := strings.Split(tagField, ",")

	keepOriginal := false
	uni := false
	res := make([]string, 0)
	pMap := make(map[string]string)
	for _, tag := range tags {
		if tag == keep {
			keepOriginal = true
			continue
		} else if tag == unique {
			uni = true
			continue
		}
		// res = append(res, tag)
		ptag := strings.ToLower(strings.Trim(strings.Split(tag, "=")[0], " "))
		pMap[ptag] = tag
		ptag = strings.ToLower(strings.Trim(strings.Split(tag, ":")[0], " "))
		pMap[ptag] = tag
	}
	// Priority
	for _, ptag := range PriorityTags {
		if tag, ok := pMap[ptag]; ok {
			if ptag == ONEOF {
				res = append(res, tags...)
			} else {
				res = append(res, tag)
			}
			delete(pMap, ptag)
		}
	}
	// custom,keep,unique
	if len(res) < 1 {
		if !keepOriginal && !uni {
			res = append(res, tags...)
		}
	}

	return structTag{
		fieldType:    strings.Join(res, ","),
		unique:       uni,
		keepOriginal: keepOriginal,
	}
}

type structTag struct {
	fieldType    string
	unique       bool
	keepOriginal bool
}

func setDataWithTag(v reflect.Value, tag string, opt options.Options) error {
	if v.Kind() != reflect.Ptr {
		return errors.New(fakerErrors.ErrValueNotPtr)
	}
	v = reflect.Indirect(v)
	switch v.Kind() {
	case reflect.Ptr:
		if _, exist := mapperTag[tag]; !exist {
			newv := reflect.New(v.Type().Elem())
			if err := setDataWithTag(newv, tag, opt); err != nil {
				return err
			}
			v.Set(newv)
			return nil
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
		return userDefinedString(v, tag, opt)
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return userDefinedNumber(v, tag)
	case reflect.Slice, reflect.Array:
		return userDefinedArray(v, tag, opt)
	case reflect.Map:
		return userDefinedMap(v, tag, opt)
	default:
		if _, exist := mapperTag[tag]; !exist {
			return fmt.Errorf(fakerErrors.ErrTagNotSupported, tag)
		}
		res, err := mapperTag[tag](v)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(res))
	}
	return nil
}

func userDefinedMap(v reflect.Value, tag string, opt options.Options) error {
	if tagFunc, ok := mapperTag[tag]; ok {
		res, err := tagFunc(v)
		if err != nil {
			return err
		}

		v.Set(reflect.ValueOf(res))
		return nil
	}

	length := randomSliceAndMapSize(opt)
	if opt.SetSliceMapNilIfLenZero && length == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	definedMap := reflect.MakeMap(v.Type())
	for i := 0; i < length; i++ {
		key, err := getValueWithTag(v.Type().Key(), tag, opt)
		if err != nil {
			return err
		}
		val, err := getValueWithTag(v.Type().Elem(), tag, opt)
		if err != nil {
			return err
		}
		definedMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	v.Set(definedMap)
	return nil
}

func getValueWithTag(t reflect.Type, tag string, opt options.Options) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		res, err := extractNumberFromTag(tag, t)
		if err != nil {
			return nil, err
		}
		return res, nil
	case reflect.String:
		res, err := extractStringFromTag(tag, opt)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return 0, errors.New(fakerErrors.ErrUnknownType)
	}
}

func getNumberWithBoundary(t reflect.Type, boundary interfaces.RandomIntegerBoundary) (interface{}, error) {
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
		return nil, errors.New(fakerErrors.ErrNotSupportedTypeForTag)
	}
}

func getValueWithNoTag(t reflect.Type, opt options.Options) (interface{}, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		boundary := interfaces.RandomIntegerBoundary{
			Start: opt.RandomIntegerBoundary.Start,
			End:   opt.RandomIntegerBoundary.End}
		res, err := getNumberWithBoundary(t, boundary)
		if err != nil {
			return nil, err
		}
		return res, nil
	case reflect.String:
		res, err := randomString(opt.RandomStringLength, opt)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return 0, errors.New(fakerErrors.ErrUnknownType)
	}
}

func userDefinedArray(v reflect.Value, tag string, opt options.Options) error {
	_, tagExists := mapperTag[tag]
	if tagExists {
		res, err := mapperTag[tag](v)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(res))
		return nil
	}
	sliceLen, err := extractSliceLengthFromTag(tag, opt)
	if err != nil {
		return err
	}

	if opt.SetSliceMapNilIfLenZero && sliceLen == 0 {
		v.Set(reflect.Zero(v.Type()))
		return nil
	}
	//remove slice_len from tag string to avoid extra logic in downstream function
	tag = findSliceLenReg.ReplaceAllString(tag, "")
	array := reflect.MakeSlice(v.Type(), sliceLen, sliceLen)
	for i := 0; i < array.Len(); i++ {
		if tag == "" {
			res, err := getValueWithNoTag(v.Type().Elem(), opt)
			if err != nil {
				return err
			}
			array.Index(i).Set(reflect.ValueOf(res))
			continue
		}

		res, err := getValueWithTag(v.Type().Elem(), tag, opt)
		if err != nil {
			return err
		}
		array.Index(i).Set(reflect.ValueOf(res))
	}
	v.Set(array)
	return nil
}

func userDefinedString(v reflect.Value, tag string, opt options.Options) error {
	var res interface{}
	var err error

	if tagFunc, ok := mapperTag[tag]; ok {
		res, err = tagFunc(v)
		if err != nil {
			return err
		}
	} else {
		res, err = extractStringFromTag(tag, opt)
		if err != nil {
			return err
		}
	}
	if res == nil {
		return fmt.Errorf(fakerErrors.ErrTagNotSupported, tag)
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
		return fmt.Errorf(fakerErrors.ErrTagNotSupported, tag)
	}

	v.Set(reflect.ValueOf(res))
	return nil
}

//extractSliceLengthFromTag checks if the sliceLength tag 'slice_len' is set, if so, returns its value, else return a random length
func extractSliceLengthFromTag(tag string, opt options.Options) (int, error) {
	if strings.Contains(tag, SliceLength) {
		lenParts := strings.SplitN(findSliceLenReg.FindString(tag), Equals, -1)
		if len(lenParts) != 2 {
			return 0, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, tag)
		}
		sliceLen, err := strconv.Atoi(lenParts[1])
		if err != nil {
			return 0, fmt.Errorf("the given sliceLength has to be numeric, tag: %s", tag)
		}
		if sliceLen < 0 {
			return 0, fmt.Errorf("slice length can not be negative, tag: %s", tag)
		}
		return sliceLen, nil
	}

	return randomSliceAndMapSize(opt), nil //Returns random slice length if the sliceLength tag isn't set
}

func extractStringFromTag(tag string, opts options.Options) (interface{}, error) {
	var err error
	strlen := opts.RandomStringLength
	strlng := opts.StringLanguage
	isOneOfTag := strings.Contains(tag, ONEOF)
	if !strings.Contains(tag, Length) && !strings.Contains(tag, Language) && !isOneOfTag {
		return nil, fmt.Errorf(fakerErrors.ErrTagNotSupported, tag)
	}
	if strings.Contains(tag, Length) {
		lenParts := strings.SplitN(findLenReg.FindString(tag), Equals, -1)
		if len(lenParts) != 2 {
			return nil, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, tag)
		}
		strlen, _ = strconv.Atoi(lenParts[1])
	}
	if strings.Contains(tag, Language) {
		strlng, err = extractLangFromTag(tag)
		if err != nil {
			return nil, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, tag)
		}
	}
	if isOneOfTag {
		var args []string
		args, err = fetchOneOfArgsFromTag(tag)
		if err != nil {
			return nil, err
		}
		toRet := args[rand.Intn(len(args))]
		return strings.TrimSpace(toRet), nil
	}

	copyOption := opts
	copyOption.StringLanguage = strlng
	res, err := randomString(strlen, copyOption)
	return res, err
}

func extractLangFromTag(tag string) (*interfaces.LangRuneBoundary, error) {
	text := findLangReg.FindString(tag)
	texts := strings.SplitN(text, Equals, -1)
	if len(texts) != 2 {
		return nil, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, text)
	}
	switch strings.ToLower(texts[1]) {
	case "eng":
		return &interfaces.LangENG, nil
	case "rus":
		return &interfaces.LangRUS, nil
	case "chi":
		return &interfaces.LangCHI, nil
	case "jpn":
		return &interfaces.LangJPN, nil
	case "kor":
		return &interfaces.LangKOR, nil
	case "emj":
		return &interfaces.EmotEMJ, nil
	default:
		return &interfaces.LangENG, nil
	}
}

func extractNumberFromTag(tag string, t reflect.Type) (interface{}, error) {
	hasOneOf := strings.Contains(tag, ONEOF)
	hasBoundaryStart := strings.Contains(tag, BoundaryStart)
	hasBoundaryEnd := strings.Contains(tag, BoundaryEnd)
	usingOneOfTag := hasOneOf && (!hasBoundaryStart && !hasBoundaryEnd)
	usingBoundariesTags := !hasOneOf && (hasBoundaryStart && hasBoundaryEnd)
	if !usingOneOfTag && !usingBoundariesTags {
		return nil, fmt.Errorf(fakerErrors.ErrTagNotSupported, tag)
	}

	// handling oneof tag
	if usingOneOfTag {
		args, err := fetchOneOfArgsFromTag(tag)
		if err != nil {
			return nil, err
		}
		switch t.Kind() {
		case reflect.Float64:
			{
				toRet, err := extractFloat64FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(float64), nil
			}
		case reflect.Float32:
			{
				toRet, err := extractFloat32FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(float32), nil
			}
		case reflect.Int64:
			{
				toRet, err := extractInt64FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(int64), nil
			}
		case reflect.Int32:
			{
				toRet, err := extractInt32FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(int32), nil
			}
		case reflect.Int16:
			{
				toRet, err := extractInt16FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(int16), nil
			}
		case reflect.Int8:
			{
				toRet, err := extractInt8FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(int8), nil
			}
		case reflect.Int:
			{
				toRet, err := extractIntFromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(int), nil
			}
		case reflect.Uint64:
			{
				toRet, err := extractUint64FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(uint64), nil
			}
		case reflect.Uint32:
			{
				toRet, err := extractUint32FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(uint32), nil
			}
		case reflect.Uint16:
			{
				toRet, err := extractUint16FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(uint16), nil
			}
		case reflect.Uint8:
			{
				toRet, err := extractUint8FromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(uint8), nil
			}
		case reflect.Uint:
			{
				toRet, err := extractUintFromTagArgs(args)
				if err != nil {
					return nil, err
				}
				return toRet.(uint), nil
			}
		default:
			{
				return nil, fmt.Errorf(fakerErrors.ErrUnsupportedNumberType)
			}
		}
	}

	// handling boundary tags
	valuesStr := strings.SplitN(tag, comma, -1)
	if len(valuesStr) != 2 {
		return nil, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, tag)
	}

	// TODO(Xaspy): When Golang provides generics, we will be able to make this method simpler and more beautiful.
	if t.Kind() == reflect.Float64 || t.Kind() == reflect.Float32 {
		startBoundary, err := extractFloatFromText(valuesStr[0])
		if err != nil {
			return nil, err
		}
		endBoundary, err := extractFloatFromText(valuesStr[1])
		if err != nil {
			return nil, err
		}
		boundary := interfaces.RandomFloatBoundary{Start: startBoundary, End: endBoundary}
		switch t.Kind() {
		case reflect.Float32:
			return float32(randomFloatWithBoundary(boundary)), nil
		case reflect.Float64:
			return randomFloatWithBoundary(boundary), nil
		}
	}

	startBoundary, err := extractIntFromText(valuesStr[0])
	if err != nil {
		return nil, err
	}
	endBoundary, err := extractIntFromText(valuesStr[1])
	if err != nil {
		return nil, err
	}
	boundary := interfaces.RandomIntegerBoundary{Start: startBoundary, End: endBoundary}
	return getNumberWithBoundary(t, boundary)
}

func extractIntFromText(text string) (int, error) {
	text = strings.TrimSpace(text)
	texts := strings.SplitN(text, Equals, -1)
	if len(texts) != 2 {
		return 0, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, text)
	}
	return strconv.Atoi(texts[1])
}

func extractFloatFromText(text string) (float64, error) {
	text = strings.TrimSpace(text)
	texts := strings.SplitN(text, Equals, -1)
	if len(texts) != 2 {
		return 0, fmt.Errorf(fakerErrors.ErrWrongFormattedTag, text)
	}
	return strconv.ParseFloat(texts[1], 64)
}

func fetchOneOfArgsFromTag(tag string) ([]string, error) {
	items := strings.Split(tag, colon)
	argsList := items[1:]
	if len(argsList) != 1 {
		return nil, fmt.Errorf(fakerErrors.ErrUnsupportedTagArguments)
	}
	if strings.Contains(argsList[0], ",,") {
		return nil, fmt.Errorf(fakerErrors.ErrDuplicateSeparator)
	}
	if argsList[0] == "" {
		return nil, fmt.Errorf(fakerErrors.ErrNotEnoughTagArguments)
	}
	args := strings.Split(argsList[0], comma)
	if len(args) < 1 {
		return nil, fmt.Errorf(fakerErrors.ErrNotEnoughTagArguments)
	}
	return args, nil
}

func randomString(n int, fakerOpt options.Options) (string, error) {
	b := make([]rune, 0)
	set := make(map[rune]struct{})
	if fakerOpt.StringLanguage.Exclude != nil {
		for _, s := range fakerOpt.StringLanguage.Exclude {
			set[s] = struct{}{}
		}
	}

	counter := 0
	for i := 0; i < n; {
		randRune := rune(rand.Intn(int(fakerOpt.StringLanguage.End-fakerOpt.StringLanguage.Start)) + int(fakerOpt.StringLanguage.Start))
		for slice.ContainsRune(set, randRune) {
			if counter++; counter >= fakerOpt.MaxGenerateStringRetries {
				return "", errors.New("Max number of string generation retries exhausted")
			}
			randRune = rune(rand.Intn(int(fakerOpt.StringLanguage.End-fakerOpt.StringLanguage.Start)) + int(fakerOpt.StringLanguage.Start))
			_, ok := set[randRune]
			if !ok {
				break
			}
		}
		b = append(b, randRune)
		i++
	}

	k := string(b)
	return k, nil
}

// randomIntegerWithBoundary returns a random integer between input start and end boundary. [start, end)
func randomIntegerWithBoundary(boundary interfaces.RandomIntegerBoundary) int {
	span := boundary.End - boundary.Start
	if span <= 0 {
		return boundary.Start
	}
	return rand.Intn(span) + boundary.Start
}

// randomFloatWithBoundary returns a random float between input start and end boundary. [start, end)
func randomFloatWithBoundary(boundary interfaces.RandomFloatBoundary) float64 {
	span := boundary.End - boundary.Start
	if span <= 0 {
		return boundary.Start
	}
	return boundary.Start + rand.Float64()*span
}

// randomInteger returns a random integer between start and end boundary. [start, end)
func randomInteger(opt *options.Options) int {
	if opt == nil {
		opt = options.DefaultOption()
	}

	return randomIntegerWithBoundary(*opt.RandomIntegerBoundary)
}

// randomFloat returns a random float between start and end boundary. [start, end)
func randomFloat(opt *options.Options) float64 {
	if opt == nil {
		opt = options.DefaultOption()
	}

	return randomFloatWithBoundary(*opt.RandomFloatBoundary)
}

// randomSliceAndMapSize returns a random integer between [0,randomSliceAndMapSize). If the testRandZero is set, returns 0
// Written for test purposes for shouldSetNil
func randomSliceAndMapSize(opt options.Options) int {
	if opt.SetSliceMapRandomToZero {
		return 0
	}
	r := opt.RandomMaxSliceSize - opt.RandomMinSliceSize
	if r < 1 {
		r = 1
	}
	return opt.RandomMinSliceSize + rand.Intn(r)
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
// (minimum_int, maximum_int, count)
// 		If only set one parameter :  An integer greater than minimum_int will be returned
// 		If only set two parameters : All integers between minimum_int and maximum_int will be returned, in a random order.
// 		If three parameters: `count` integers between minimum_int and maximum_int will be returned.
func RandomInt(parameters ...int) (p []int, err error) {
	switch len(parameters) {
	case 1:
		minInt := parameters[0]
		p = rand.Perm(minInt)
		for i := range p {
			p[i] += minInt
		}
	case 2:
		minInt, maxInt := parameters[0], parameters[1]
		p = rand.Perm(maxInt - minInt + 1)

		for i := range p {
			p[i] += minInt
		}
	case 3:
		minInt, maxInt := parameters[0], parameters[1]
		count := parameters[2]
		p = rand.Perm(maxInt - minInt + 1)

		for i := range p {
			p[i] += minInt
		}
		if len(p) > count {
			p = p[0:count]
		}
	default:
		err = fmt.Errorf(fakerErrors.ErrMoreArguments, len(parameters))
	}
	return p, err
}

func generateUnique(dataType string, fn func() interface{}) (interface{}, error) {
	for i := 0; i < maxRetry; i++ {
		value := fn()
		if !slice.ContainsValue(uniqueValues[dataType], value) { // Retry if unique value already found
			uniqueValues[dataType] = append(uniqueValues[dataType], value)
			return value, nil
		}
	}
	return reflect.Value{}, fmt.Errorf(fakerErrors.ErrUniqueFailure, dataType)
}

func singleFakeData(dataType string, fn func() interface{}, opts ...options.OptionFunc) interface{} {
	ops := initOption(opts...)
	if ops.GenerateUniqueValues {
		v, err := generateUnique(dataType, fn)
		if err != nil {
			panic(err)
		}
		return v
	}
	return fn()
}
