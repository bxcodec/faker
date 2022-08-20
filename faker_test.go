package faker

import (
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	fakerErrors "github.com/bxcodec/faker/v4/pkg/errors"
	"github.com/bxcodec/faker/v4/pkg/interfaces"
	"github.com/bxcodec/faker/v4/pkg/options"
)

const (
	someStructLen           = 2
	someStructBoundaryStart = 5
	someStructBoundaryEnd   = 10

	someStructWithLenAndLangENG = 5
	someStructWithLenAndLangCHI = 10
	someStructWithLenAndLangRUS = 15
	someStructWithLenAndLangJPN = 20
	someStructWithLenAndLangKOR = 25
	someStructWithLenAndEmotEMJ = 50
)

var (
	langCorrectTagsMap = map[string]interfaces.LangRuneBoundary{"lang=eng": interfaces.LangENG,
		"lang=chi": interfaces.LangCHI, "lang=rus": interfaces.LangRUS, "lang=jpn": interfaces.LangJPN,
		"lang=kor": interfaces.LangKOR, "lang=emj": interfaces.EmotEMJ}
	langUncorrectTags = [3]string{"lang=", "lang", "lng=eng"}

	lenCorrectTags   = [3]string{"len=4", "len=5", "len=10"}
	lenUncorrectTags = [6]string{"len=b", "ln=10", "length=25", "lang=b", "ln=10", "lang=8d,,len=eng"}

	sliceLenCorrectTags   = [4]string{"slice_len=0", "slice_len=4", "slice_len=5", "slice_len=10"}
	sliceLenIncorrectTags = [3]string{"slice_len=b", "slice_len=-1", "slice_len=-10"}
)

type Coupon struct {
	ID         int      `json:"id" xorm:"id"`
	BrokerCode string   `json:"broker_code" xorm:"broker_code"`
	IgetUID    int      `json:"iget_uid" xorm:"iget_uid"`
	CreateTime string   `json:"create_time" xorm:"create_time"`
	CFirstName string   `json:"chinese_first_name" faker:"chinese_first_name"`
	CLsstName  string   `json:"chinese_last_name" faker:"chinese_last_name"`
	CName      string   `json:"name" faker:"chinese_name"`
	AdNames    []string `json:"ad_name" xorm:"ad_name" faker:"slice_len=5,len=10"` // faker:"len=10,slice_len=5"
	CdNames    []string `json:"cd_name" xorm:"cd_name" faker:"len=10,slice_len=5"` //
}

func TestPLen(t *testing.T) {
	coupon := Coupon{}
	err := FakeData(&coupon)
	if err != nil {
		t.Fatal(err)
		return
	}
	if len(coupon.AdNames[0]) != 10 || len(coupon.AdNames) != 5 {
		t.Fatal("slice len is error")
	}
	if len(coupon.CdNames[0]) != 10 || len(coupon.CdNames) != 5 {
		t.Fatal("slice len is error")
	}
	t.Logf("%+v\n", coupon)
}

type SomeInt32 int32

type TArray [16]byte

type SomeStruct struct {
	Inta    int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Float32 float32
	Float64 float64

	UInta  uint
	UInt8  uint8
	UInt16 uint16
	UInt32 uint32
	UInt64 uint64

	Latitude           float32 `faker:"lat"`
	LATITUDE           float64 `faker:"lat"`
	Long               float32 `faker:"long"`
	LONG               float64 `faker:"long"`
	StringValue        string
	CreditCardType     string `faker:"cc_type"`
	CreditCardNumber   string `faker:"cc_number"`
	Email              string `faker:"email"`
	IPV4               string `faker:"ipv4"`
	IPV6               string `faker:"ipv6"`
	Bool               bool
	SString            []string
	SInt               []int
	SInt8              []int8
	SInt16             []int16
	SInt32             []int32
	SInt64             []int64
	SFloat32           []float32
	SFloat64           []float64
	SBool              []bool
	Struct             AStruct
	TArray             TArray
	Time               time.Time
	Stime              []time.Time
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	ID                 string  `faker:"uuid_digit"`
	HyphenatedID       string  `faker:"uuid_hyphenated"`

	MapStringString        map[string]string
	MapStringStruct        map[string]AStruct
	MapStringStructPointer map[string]*AStruct

	SomeInt32s []SomeInt32
}

type SomeStructWithLen struct {
	Inta  int   `faker:"boundary_start=5, boundary_end=10"`
	Int8  int8  `faker:"boundary_start=5, boundary_end=10"`
	Int16 int16 `faker:"boundary_start=5, boundary_end=10"`
	Int32 int32 `faker:"boundary_start=5, boundary_end=10"`
	Int64 int64 `faker:"boundary_start=5, boundary_end=10"`

	UInta  uint   `faker:"boundary_start=5, boundary_end=10"`
	UInt8  uint8  `faker:"boundary_start=5, boundary_end=10"`
	UInt16 uint16 `faker:"boundary_start=5, boundary_end=10"`
	UInt32 uint32 `faker:"boundary_start=5, boundary_end=10"`
	UInt64 uint64 `faker:"boundary_start=5, boundary_end=10"`

	Float32 float32 `faker:"boundary_start=5, boundary_end=10"`
	Float64 float64 `faker:"boundary_start=5, boundary_end=10"`

	ASString []string          `faker:"len=2"`
	SString  string            `faker:"len=2"`
	MSString map[string]string `faker:"len=2"`
	MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
}

type SomeStructWithLang struct {
	ValueENG string `faker:"lang=eng"`
	ValueCHI string `faker:"lang=chi"`
	ValueRUS string `faker:"lang=rus"`
	ValueJPN string `faker:"lang=jpn"`
	ValueKOR string `faker:"lang=kor"`
	ValueEMJ string `faker:"lang=emj"`

	ValueWithUndefinedLang string `faker:"lang=und"`
}

type SomeStructWithLenAndLang struct {
	ValueENG string `faker:"len=5, lang=eng"`
	ValueCHI string ` faker:"len=10, lang=chi"`
	ValueRUS string ` faker:"len=15, lang=rus"`
	ValueJPN string ` faker:"len=20, lang=jpn"`
	ValueKOR string ` faker:"len=25, lang=kor"`
	ValueEMJ string ` faker:"len=50, lang=emj"`
}

func (s SomeStruct) String() string {
	return fmt.Sprintf(`{
	Inta: %v
	Int8: %v
	Int16: %v
	Int32: %v
	Int64: %v
	Float32: %v
	Float64: %v

	UInta: %v
	UInt8: %v
	UInt16: %v
	UInt32: %v
	UInt64: %v

	Latitude: %v
	LATITUDE: %v
	Long: %v
	LONG: %v
	StringValue: %v
	CreditCardType: %v
	CreditCardNumber: %v
	Email: %v
	IPV4: %v
	IPV6: %v
	Bool: %v
	SString: %v
	SInt: %v
	SInt8: %v
	SInt16: %v
	SInt32: %v
	SInt64: %v
	SFloat32: %v
	SFloat64:%v
	SBool: %v
	Struct: %v
	Time: %v
	Stime: %v
	Currency: %v
	Amount: %v
	AmountWithCurrency: %v
	ID: %v
	HyphenatedID: %v

	MapStringString: %v
	MapStringStruct: %v
	MapStringStructPointer: %v
	}`, s.Inta, s.Int8, s.Int16, s.Int32,
		s.Int64, s.Float32, s.Float64, s.UInta,
		s.UInt8, s.UInt16, s.UInt32, s.UInt64,
		s.Latitude, s.LATITUDE, s.Long, s.LONG,
		s.StringValue, s.CreditCardType, s.CreditCardNumber,
		s.Email, s.IPV4, s.IPV6, s.Bool, s.SString, s.SInt,
		s.SInt8, s.SInt16, s.SInt32, s.SInt64, s.SFloat32, s.SFloat64,
		s.SBool, s.Struct, s.Time, s.Stime, s.Currency, s.Amount,
		s.AmountWithCurrency, s.ID, s.HyphenatedID, s.MapStringString,
		s.MapStringStruct, s.MapStringStructPointer)
}

type AStruct struct {
	Number        int64
	Height        int64
	AnotherStruct CStruct
}

type BStruct struct {
	Image string
}
type CStruct struct {
	BStruct
	Name string
}

type TaggedStruct struct {
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	Jwt                string  `faker:"jwt"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	ChineseFirstName   string  `faker:"chinese_first_name"`
	ChineseLastName    string  `faker:"chinese_last_name"`
	ChineseName        string  `faker:"chinese_name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float32 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	ID                 string  `faker:"uuid_digit"`
	HyphenatedID       string  `faker:"uuid_hyphenated"`
}

func (t TaggedStruct) String() string {
	return fmt.Sprintf(`{
	Latitude: %f,
	Long: %f,
	CreditCardNumber: %s,
	CreditCardType: %s,
	Email: %s,
	DomainName: %s,
	IPV4: %s,
	IPV6: %s,
	Password: %s,
	Jwt: %s,
	PhoneNumber: %s,
	MacAddress: %s,
	URL: %s,
	UserName: %s,
	TollFreeNumber: %s,
	E164PhoneNumber: %s,
	TitleMale: %s,
	TitleFemale: %s,
	FirstName: %s,
	FirstNameMale: %s,
	FirstNameFemale: %s,
	LastName: %s,
	Name: %s,
	UnixTime: %d,
	Date: %s,
	Time: %s,
	MonthName: %s,
	Year: %s,
	DayOfWeek: %s,
	DayOfMonth: %s,
	Timestamp: %s,
	Century: %s,
	TimeZone: %s,
	TimePeriod: %s,
	Word: %s,
	Sentence: %s,
	Paragraph: %s,
	Currency: %s,
	Amount: %f,
	AmountWithCurrency: %s,
	HyphenatedID: %s,
	ID: %s,
}`, t.Latitude, t.Longitude, t.CreditCardNumber,
		t.CreditCardType, t.Email, t.DomainName, t.IPV4,
		t.IPV6, t.Password, t.Jwt, t.PhoneNumber, t.MacAddress,
		t.URL, t.UserName, t.TollFreeNumber,
		t.E164PhoneNumber, t.TitleMale, t.TitleFemale,
		t.FirstName, t.FirstNameMale, t.FirstNameFemale, t.LastName,
		t.Name, t.UnixTime, t.Date,
		t.Time, t.MonthName, t.Year, t.DayOfWeek,
		t.DayOfMonth, t.Timestamp, t.Century, t.TimeZone,
		t.TimePeriod, t.Word, t.Sentence, t.Paragraph,
		t.Currency, t.Amount, t.AmountWithCurrency,
		t.HyphenatedID, t.ID,
	)
}

type NotTaggedStruct struct {
	Latitude         float32
	Long             float32
	CreditCardType   string
	CreditCardNumber string
	Email            string
	IPV4             string
	IPV6             string
}

func TestFakerData(t *testing.T) {
	var a SomeStruct
	err := FakeData(&a)

	if err != nil {
		t.Error("Expected NoError")
	}
	fmt.Println("SomeStruct:")
	t.Logf("%+v\n", a)

	var b TaggedStruct
	err = FakeData(&b)

	if err != nil {
		t.Error("Expected NoError, but Got Err: ", err)
	}

	fmt.Println("TaggedStruct:")
	t.Logf("%+v\n", b)

	// Example Result :
	// {Int:8906957488773767119 Int8:6 Int16:14 Int32:391219825 Int64:2374447092794071106 String:poraKzAxVbWVkMkpcZCcWlYMd Bool:false SString:[MehdV aVotHsi] SInt:[528955241289647236 7620047312653801973 2774096449863851732] SInt8:[122 -92 -92] SInt16:[15679 -19444 -30246] SInt32:[1146660378 946021799 852909987] SInt64:[6079203475736033758 6913211867841842836 3269201978513619428] SFloat32:[0.019562425 0.12729558 0.36450312] SFloat64:[0.7825838989890364 0.9732903338838912 0.8316541489234004] SBool:[true false true] Struct:{Number:7693944638490551161 Height:6513508020379591917}}

}

func TestCustomFakerOnUnsupportedMapStringInterface(t *testing.T) {
	type Sample struct {
		Map map[string]interface{} `faker:"custom"`
	}

	err := AddProvider("custom", func(v reflect.Value) (interface{}, error) {
		return map[string]interface{}{"foo": "bar"}, nil
	})
	if err != nil {
		t.Error("Expected NoError, but Got Err", err)
	}

	var sample = new(Sample)
	err = FakeData(sample)
	if err != nil {
		t.Error("Expected NoError, but Got Err:", err)
	}

	actual, ok := sample.Map["foo"]
	if !ok {
		t.Error("map key not set by custom faker")
	}

	if actual != "bar" {
		t.Error("map value not set by custom faker")
	}
}

func TestUnsuportedMapStringInterface(t *testing.T) {
	type Sample struct {
		Map map[string]interface{}
	}
	var sample = new(Sample)
	if err := FakeData(sample); err == nil {
		t.Error("Expected Error. But got nil")
	}
}

func TestSetDataIfArgumentNotPtr(t *testing.T) {
	temp := struct{}{}
	if "Not a pointer value" != FakeData(temp).Error() {
		t.Error("Expected in arguments not ptr")
	}
}

func TestSetDataIfArgumentNotHaveReflect(t *testing.T) {
	temp := func() {}

	if err := FakeData(temp); err == nil {
		t.Error("Exptected error but got nil")
	}
}

func TestSetDataErrorDataParseTagStringType(t *testing.T) {
	temp := &struct {
		Test string `faker:"test_no_exist"`
	}{}
	for idx, tag := range PriorityTags {
		if tag == "test_no_exist" {
			PriorityTags[idx] = ""
		}
	}
	if err := FakeData(temp); err == nil {
		t.Error("Exptected error Unsupported tag, but got nil", temp, err)
	}

}

func TestSetDataErrorDataParseTagIntType(t *testing.T) {
	temp := &struct {
		Test int `faker:"test_no_exist"`
	}{}
	for idx, tag := range PriorityTags {
		if tag == "test_no_exist" {
			PriorityTags[idx] = ""
		}
	}
	if err := FakeData(temp); err == nil {
		t.Error("Expected error Unsupported tag, but got nil")
	}
}

func TestSetRandomStringLength(t *testing.T) {
	someStruct := SomeStruct{}
	strLen := 5
	if err := FakeData(&someStruct, options.WithRandomStringLength(uint(strLen))); err != nil {
		t.Error("Fake data generation has failed")
	}
	if utfLen(someStruct.StringValue) > strLen {
		t.Error("SetRandomStringLength did not work.")
	}
}

func TestSetStringLang(t *testing.T) {
	someStruct := SomeStruct{}
	// optionsSetStringLang(LangENG)
	if err := FakeData(&someStruct, options.WithStringLanguage(interfaces.LangENG)); err != nil {
		t.Error("Fake data generation has failed")
	}
}

func TestSetRandomNumberBoundaries(t *testing.T) {
	someStruct := SomeStruct{}
	boundary := interfaces.RandomIntegerBoundary{Start: 10, End: 90}
	if err := FakeData(&someStruct, options.WithRandomIntegerBoundaries(boundary)); err != nil {
		t.Error("Fake data generation has failed")
	}

	if someStruct.Inta >= boundary.End || someStruct.Inta < boundary.Start {
		t.Errorf("%d must be between [%d,%d)", someStruct.Inta, boundary.Start, boundary.End)
	}
}

func TestSetRandomMapAndSliceSize(t *testing.T) {
	someStruct := SomeStruct{}
	size := 5
	if err := FakeData(&someStruct, options.WithRandomMapAndSliceMaxSize(uint(size))); err != nil {
		t.Error("Fake data generation has failed")
	}
	if len(someStruct.MapStringStruct) > size || len(someStruct.SBool) > size {
		t.Error("SetRandomMapAndSliceSize did not work.")
	}
}

func TestSetNilIfLenIsZero(t *testing.T) {
	someStruct := SomeStruct{}
	// testRandZero = true
	if err := FakeData(&someStruct, options.WithNilIfLenIsZero(true), options.WithSliceMapRandomToZero(true)); err != nil {
		t.Error("Fake data generation has failed")
	}
	if someStruct.MapStringString != nil && someStruct.MapStringStruct != nil &&
		someStruct.MapStringStructPointer != nil {
		t.Error("Map has to be nil")
	}
	if someStruct.Stime != nil && someStruct.SBool != nil {
		t.Error("Array has to be nil")
	}
}

func TestSetIgnoreInterface(t *testing.T) {
	var someInterface interface{}
	if err := FakeData(&someInterface, options.WithIgnoreInterface(false)); err == nil {
		t.Error("Fake data generation didn't fail on interface{}")
	}
	if err := FakeData(&someInterface, options.WithIgnoreInterface(true)); err != nil {
		t.Error("Fake data generation fail on interface{} with SetIgnoreInterface(true)")
	}
}

func TestBoundaryAndLen(t *testing.T) {
	iterate := 10
	someStruct := SomeStructWithLen{}
	for i := 0; i < iterate; i++ {
		if err := FakeData(&someStruct); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.Int8)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.Int16)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.Int32)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(someStruct.Inta); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.Int64)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.UInt8)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.UInt16)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.UInt32)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.UInta)); err != nil {
			t.Error(err)
		}
		if err := validateIntRange(int(someStruct.UInt64)); err != nil {
			t.Error(err)
		}
		if err := validateFloatRange(float64(someStruct.Float32)); err != nil {
			t.Error(err)
		}
		if err := validateFloatRange(someStruct.Float64); err != nil {
			t.Error(err)
		}
		if err := validateLen(someStruct.SString); err != nil {
			t.Error(err)
		}
		for _, str := range someStruct.ASString {
			if err := validateLen(str); err != nil {
				t.Error(err)
			}
		}
		for k, v := range someStruct.MSString {
			if err := validateLen(k); err != nil {
				t.Error(err)
			}
			if err := validateLen(v); err != nil {
				t.Error(err)
			}
		}
		for k, v := range someStruct.MIint {
			if err := validateIntRange(k); err != nil {
				t.Error(err)
			}
			if err := validateIntRange(v); err != nil {
				t.Error(err)
			}
		}
	}
}

func TestWrongBoundaryAndLen(t *testing.T) {
	type SomeStruct struct {
		Value int `faker:"boundary_start=a, boundary_end=b"`
	}
	s := SomeStruct{}
	if err := FakeData(&s); err == nil {
		t.Error(err)
	}
}

func TestLang(t *testing.T) {
	someStruct := SomeStructWithLang{}
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}

	var err error
	err = isStringLangCorrect(someStruct.ValueENG, interfaces.LangENG)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(someStruct.ValueRUS, interfaces.LangRUS)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(someStruct.ValueCHI, interfaces.LangCHI)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(someStruct.ValueJPN, interfaces.LangJPN)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(someStruct.ValueKOR, interfaces.LangKOR)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(someStruct.ValueEMJ, interfaces.EmotEMJ)
	if err != nil {
		t.Error(err.Error())
	}

	err = isStringLangCorrect(someStruct.ValueWithUndefinedLang, interfaces.LangENG)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLangWithWrongLang(t *testing.T) {
	type SomeStruct struct {
		String string `faker:"lang=undefined"`
	}

	s := SomeStruct{}
	err := FakeData(&s)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestExtractingLangFromTag(t *testing.T) {
	var err error
	var lng *interfaces.LangRuneBoundary
	for k, v := range langCorrectTagsMap {
		if lng, err = extractLangFromTag(k); err != nil {
			t.Error(err.Error())
		}
		if !reflect.DeepEqual(v, *lng) {
			t.Errorf("Got %v lang rune range, but expected %v", lng, k)
		}
	}
	for _, tag := range langUncorrectTags {
		if _, err := extractLangFromTag(tag); err == nil {
			t.Error(err.Error())
		}
	}

}

func TestExtractingStringFromTag(t *testing.T) {
	for _, tag := range lenCorrectTags {
		if _, err := extractStringFromTag(tag, *options.DefaultOption()); err != nil {
			t.Error(err.Error())
		}
	}
	for _, tag := range lenUncorrectTags {
		if _, err := extractStringFromTag(tag, *options.DefaultOption()); err == nil {
			t.Error(err.Error())
		}
	}
}

func TestSliceLen(t *testing.T) {
	type SomeStruct struct {
		String1 []string `faker:"slice_len=0"`
		String2 []string `faker:"slice_len=5"`
		String3 []string `faker:"slice_len=10"`
	}
	var someStruct SomeStruct
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}

	if len(someStruct.String1) != 0 {
		t.Errorf("Wrong slice length based on slice_len tag, got %d, wanted 0", len(someStruct.String1))
	}
	if len(someStruct.String2) != 5 {
		t.Errorf("Wrong slice length based on slice_len tag, got %d, wanted 5", len(someStruct.String2))
	}
	if len(someStruct.String3) != 10 {
		t.Errorf("Wrong slice length based on slice_len tag, got %d, wanted 10", len(someStruct.String3))
	}
}

func TestWrongSliceLen(t *testing.T) {
	type SomeStruct struct {
		String []string `faker:"slice_len=bla"`
	}

	s := SomeStruct{}
	err := FakeData(&s)

	if err == nil {
		t.Error("An error should be thrown for the wrong slice_len")
	}
}

func TestExtractingSliceLenFromTag(t *testing.T) {
	for _, tag := range sliceLenCorrectTags {
		if _, err := extractSliceLengthFromTag(tag, *options.DefaultOption()); err != nil {
			t.Error(err.Error())
		}
	}
	for _, tag := range sliceLenIncorrectTags {
		if _, err := extractSliceLengthFromTag(tag, *options.DefaultOption()); err == nil {
			t.Errorf("Extracting should have thrown an error for tag %s", tag)
		}
	}
}

func TestLangWithLen(t *testing.T) {
	someStruct := SomeStructWithLenAndLang{}
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}

	var err error
	err = isStringLangCorrect(someStruct.ValueENG, interfaces.LangENG)
	if err != nil {
		t.Error(err.Error())
	}
	engLen := utfLen(someStruct.ValueENG)
	if engLen != someStructWithLenAndLangENG {
		t.Errorf("Got %d, but expected to be %d as a string len", engLen, someStructWithLenAndLangENG)
	}

	err = isStringLangCorrect(someStruct.ValueRUS, interfaces.LangRUS)
	if err != nil {
		t.Error(err.Error())
	}
	chiLen := utfLen(someStruct.ValueCHI)
	if chiLen != someStructWithLenAndLangCHI {
		t.Errorf("Got %d, but expected to be %d as a string len", chiLen, someStructWithLenAndLangCHI)
	}

	err = isStringLangCorrect(someStruct.ValueCHI, interfaces.LangCHI)
	if err != nil {
		t.Error(err.Error())
	}
	rusLen := utfLen(someStruct.ValueRUS)
	if rusLen != someStructWithLenAndLangRUS {
		t.Errorf("Got %d, but expected to be %d as a string len", rusLen, someStructWithLenAndLangRUS)
	}

	err = isStringLangCorrect(someStruct.ValueJPN, interfaces.LangJPN)
	if err != nil {
		t.Error(err.Error())
	}
	jpnLen := utfLen(someStruct.ValueJPN)
	if jpnLen != someStructWithLenAndLangJPN {
		t.Errorf("Got %d, but expected to be %d as a string len", jpnLen, someStructWithLenAndLangJPN)
	}

	err = isStringLangCorrect(someStruct.ValueKOR, interfaces.LangKOR)
	if err != nil {
		t.Error(err.Error())
	}
	korLen := utfLen(someStruct.ValueKOR)
	if korLen != someStructWithLenAndLangKOR {
		t.Errorf("Got %d, but expected to be %d as a string len", korLen, someStructWithLenAndLangKOR)
	}

	err = isStringLangCorrect(someStruct.ValueEMJ, interfaces.EmotEMJ)
	if err != nil {
		t.Error(err.Error())
	}
	emjLen := utfLen(someStruct.ValueEMJ)
	if emjLen != someStructWithLenAndEmotEMJ {
		t.Errorf("Got %d, but expected to be %d as a string len", emjLen, someStructWithLenAndEmotEMJ)
	}
}

func isStringLangCorrect(value string, lang interfaces.LangRuneBoundary) error {
	for i := 0; i < len(value); {
		r, size := utf8.DecodeLastRuneInString(value[i:])
		if r < lang.Start || r > lang.End {
			return fmt.Errorf("Symbol is not in selected alphabet: start: %d, end: %d", lang.Start, lang.End)
		}
		i += size
	}
	return nil
}

func TestExtractNumberFromTagFail(t *testing.T) {
	notSupportedStruct := &struct {
		Test int `faker:"boundary_start=5"`
	}{}
	if err := FakeData(&notSupportedStruct); err == nil {
		t.Error(err)
	}
	wrongFormatStruct := &struct {
		Test int `faker:"boundary_start=5 boundary_end=10"`
	}{}
	if err := FakeData(&wrongFormatStruct); err == nil {
		t.Error(err)
	}
	startExtractionStruct := &struct {
		Test int `faker:"boundary_start=asda, boundary_end=10"`
	}{}
	if err := FakeData(&startExtractionStruct); err == nil {
		t.Error(err)
	}
	endExtractionStruct := &struct {
		Test int `faker:"boundary_start=5, boundary_end=asda"`
	}{}
	if err := FakeData(&endExtractionStruct); err == nil {
		t.Error(err)
	}
	wrongSplitFormatStruct := &struct {
		Test int `faker:"boundary_start5, boundary_end=10"`
	}{}
	if err := FakeData(&wrongSplitFormatStruct); err == nil {
		t.Error(err)
	}
}

func TestUserDefinedStringFail(t *testing.T) {
	wrongFormatStruct := &struct {
		Test string `faker:"len=asd"`
	}{}
	if err := FakeData(&wrongFormatStruct); err == nil {
		t.Error(err)
	}
}

func validateLen(value string) error {
	if len(value) != someStructLen {
		return fmt.Errorf("Got %d, but expected to be %d as a string len", len(value), someStructLen)
	}
	return nil
}

func validateIntRange(value int) error {
	if value < someStructBoundaryStart || value > someStructBoundaryEnd {
		return fmt.Errorf("%d must be between %d and %d", value, someStructBoundaryStart,
			someStructBoundaryEnd)
	}
	return nil
}

func validateFloatRange(value float64) error {
	if value < someStructBoundaryStart || value > someStructBoundaryEnd {
		return fmt.Errorf("%f must be between %d and %d", value, someStructBoundaryStart,
			someStructBoundaryEnd)
	}
	return nil
}

func TestSetDataWithTagIfFirstArgumentNotPtr(t *testing.T) {
	temp := struct{}{}
	if setDataWithTag(reflect.ValueOf(temp), "", *options.DefaultOption()).Error() != "Not a pointer value" {
		t.Error("Expected in arguments not ptr")
	}
}

func BenchmarkFakerDataNOTTagged(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := NotTaggedStruct{}
		err := FakeData(&a)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFakerDataTagged(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := TaggedStruct{}
		err := FakeData(&a)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestRandomIntOnlyFirstParameter(t *testing.T) {
	r := rand.Intn(100)
	res, _ := RandomInt(r)
	if len(res) != r {
		t.Error("It is expected that the refund amount is equal to the argument (RandomInt)")
	}
}

func TestRandomIntOnlySecondParameters(t *testing.T) {
	first := rand.Intn(50)
	second := rand.Intn(100) + first
	res, _ := RandomInt(first, second)
	if len(res) != (second - first + 1) {
		t.Error("It is expected that the refund amount is equal to the argument (RandomInt)")
	}
}

func TestRandomIntThreeParameters(t *testing.T) {
	first := rand.Intn(50)
	second := rand.Intn(100) + first
	third := rand.Intn(5)
	res, _ := RandomInt(first, second, third)
	if len(res) != (third) {
		t.Errorf("Incorrect number of results returned. Expected %v. Got %v.", third, len(res))
	}

	for _, v := range res {
		if v < first {
			t.Errorf("Found value %v below minimum %v.", v, first)
		}
		if v > second {
			t.Errorf("Found value %v above maximum %v.", v, second)
		}
	}
}

func TestRandomIntOnlyError(t *testing.T) {
	arguments := []int{1, 3, 4, 5, 6}
	_, err := RandomInt(arguments...)
	if err == nil && err.Error() == fmt.Errorf(fakerErrors.ErrMoreArguments, len(arguments)).Error() {
		t.Error("Expected error from function RandomInt")
	}
}

type PointerStructA struct {
	SomeStruct *SomeStruct
}
type PointerStructB struct {
	PointA PointerStructA
}

type PointerC struct {
	TaggedStruct *TaggedStruct
}

func TestStructPointer(t *testing.T) {
	a := new(PointerStructB)
	err := FakeData(a)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	t.Logf(" A value: %+v , Somestruct Value: %+v  ", a, a)

	tagged := new(PointerC)
	err = FakeData(tagged)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	t.Logf(" tagged value: %+v , TaggedStruct Value: %+v  ", a, a.PointA.SomeStruct)
}

type CustomString string
type CustomInt int
type CustomMap map[string]string
type CustomPointerStruct PointerStructB
type CustomTypeStruct struct {
	CustomString        CustomString
	CustomInt           CustomInt
	CustomMap           CustomMap
	CustomPointerStruct CustomPointerStruct
}

func TestCustomType(t *testing.T) {
	a := new(CustomTypeStruct)
	err := FakeData(a)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	t.Logf(" A value: %+v , Somestruct Value: %+v  ", a, a)
}

type SampleStruct struct {
	name string
	Age  int
}

func (s SampleStruct) GetName() string {
	return s.name
}

func TestUnexportedFieldStruct(t *testing.T) {
	// This test is to ensure that the faker won't panic if trying to fake data on struct that has unexported field
	a := new(SampleStruct)
	err := FakeData(a)

	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
		t.FailNow()
	}
	t.Logf(" A value: %+v , SampleStruct Value: %+v  ", a, a)
}

func TestPointerToCustomScalar(t *testing.T) {
	// This test is to ensure that the faker won't panic if trying to fake data on struct that has field
	a := new(CustomInt)
	err := FakeData(a)

	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	t.Logf(" A value: %+v , Custom scalar Value: %+v  ", a, a)
}

type PointerCustomIntStruct struct {
	V *CustomInt
}

func TestPointerToCustomIntStruct(t *testing.T) {
	// This test is to ensure that the faker won't panic if trying to fake data on struct that has field
	a := new(PointerCustomIntStruct)
	err := FakeData(a)

	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	t.Logf(" A value: %+v , PointerCustomIntStruct scalar Value: %+v  ", a, a)
}

func TestSkipField(t *testing.T) {
	// This test is to ensure that the faker won't fill field with tag skip

	a := struct {
		ID                    int
		ShouldBeSkipped       int `faker:"-"`
		ShouldBeSkippedFilled int `faker:"-"`
	}{}

	a.ShouldBeSkippedFilled = 10

	err := FakeData(&a)

	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}

	if a.ShouldBeSkipped != 0 {
		t.Error("Expected that field will be skipped")
	}

	if a.ShouldBeSkippedFilled != 10 {
		t.Error("Expected that field will be skipped")
	}
}

type Student struct {
	Name   string
	School School `faker:"custom-school"`
}
type School struct {
	Location string
}

func TestExtend(t *testing.T) {
	// This test is to ensure that faker can be extended new providers

	t.Run("test-string", func(t *testing.T) {
		a := struct {
			ID string `faker:"test"`
		}{}

		err := AddProvider("test", func(v reflect.Value) (interface{}, error) {
			return "test", nil
		})

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		err = FakeData(&a)

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		if a.ID != "test" {
			t.Error("ID should be equal test value")
		}
	})

	t.Run("test-struct", func(t *testing.T) {
		a := &Student{}
		err := AddProvider("custom-school", func(v reflect.Value) (interface{}, error) {

			sch := School{
				Location: "North Kindom",
			}

			return sch, nil
		})

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		err = FakeData(&a)

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		if a.School.Location != "North Kindom" {
			t.Error("ID should be equal test value")
		}
	})

	type CustomTypeOverSlice []byte
	type CustomThatUsesSlice struct {
		UUID CustomTypeOverSlice `faker:"custom-type-over-slice"`
	}

	t.Run("test-with-custom-slice-type", func(t *testing.T) {
		a := CustomThatUsesSlice{}
		err := AddProvider("custom-type-over-slice", func(v reflect.Value) (interface{}, error) {
			return CustomTypeOverSlice{0, 1, 2, 3, 4}, nil
		})

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		err = FakeData(&a)

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		if reflect.DeepEqual(a.UUID, []byte{0, 1, 2, 3, 4}) {
			t.Error("UUID should equal test value")
		}
	})

	type MyInt int
	type Sample struct {
		Value []MyInt `faker:"myint"`
	}

	t.Run("test with type alias for int", func(t *testing.T) {
		a := Sample{}
		sliceLen := 10
		err := AddProvider("myint", func(v reflect.Value) (interface{}, error) {
			r1 := mathrand.New(NewSafeSource(mathrand.NewSource(time.Now().UnixNano())))
			r := make([]MyInt, sliceLen)
			for i := range r {
				r[i] = MyInt(r1.Intn(100))
			}
			return r, nil
		})

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		err = FakeData(&a)

		if err != nil {
			t.Error("Expected Not Error, But Got: ", err)
		}

		if len(a.Value) != sliceLen {
			t.Errorf("Expected a slice of length %v but got %v", sliceLen, len(a.Value))
		}

	})

}

func TestTagAlreadyExists(t *testing.T) {
	// This test is to ensure that existing tag cannot be rewritten

	err := AddProvider(EmailTag, func(v reflect.Value) (interface{}, error) {
		return nil, nil
	})

	if err == nil || err.Error() != fakerErrors.ErrTagAlreadyExists {
		t.Error("Expected ErrTagAlreadyExists Error,  But Got: ", err)
	}
}

func TestRemoveProvider(t *testing.T) {
	err := AddProvider("new_test_tag", func(v reflect.Value) (interface{}, error) {
		return "test", nil
	})
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}

	err = RemoveProvider("new_test_tag")
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
}

func TestTagDoesNotExist(t *testing.T) {
	err := RemoveProvider("not_existing_test_tag")

	if err == nil || err.Error() != fakerErrors.ErrTagDoesNotExist {
		t.Error("Expected ErrTagDoesNotExist Error,  But Got: ", err)
	}
}

func TestTagWithPointer(t *testing.T) {

	type TestStruct struct {
		FirstName  *string  `json:"first_name,omitempty" faker:"first_name_male"`
		Email      *string  `json:"email,omitempty" faker:"email"`
		Latitude   *float64 `faker:"lat"`
		Latitude32 *float32 `faker:"lat"`
		UnixTime   *int64   `faker:"unix_time"`
		School     *School  `faker:"school"`
	}
	// With custom provider
	err := AddProvider("school", func(v reflect.Value) (interface{}, error) {
		return &School{Location: "Jakarta"}, nil
	})
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	var sample TestStruct
	err = FakeData(&sample)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}

	// Assert
	if sample.FirstName == nil || *sample.FirstName == "" {
		t.Error("Expected filled but got empty")
	}
	if sample.Email == nil || *sample.Email == "" {
		t.Error("Expected filled but got empty")
	}
	if sample.Latitude == nil || *sample.Latitude == 0 {
		t.Error("Expected filled but got empty")
	}
	if sample.Latitude32 == nil || *sample.Latitude32 == 0 {
		t.Error("Expected filled but got empty")
	}

	if sample.UnixTime == nil || *sample.UnixTime == 0 {
		t.Error("Expected filled but got empty")
	}

	if sample.School == nil || sample.School.Location == "" {
		t.Error("Expected filled but got empty")
	}
}

func TestItOverwritesDefaultValueIfKeepIsSet(t *testing.T) {
	type TestStruct struct {
		Email string `json:"email,omitempty" faker:"email,keep"`
	}

	test := TestStruct{}

	err := FakeData(&test)
	if err != nil {
		t.Error("expected not error, but got: ", err)
	}

	if test.Email == "" {
		t.Error("expected filled but got empty")
	}
}
func TestItKeepsStructPropertyWhenTagKeepIsSet(t *testing.T) {
	type TestStruct struct {
		FirstName string            `json:"first_name,omitempty" faker:"first_name_male,keep"`
		Email     string            `json:"email,omitempty" faker:"email,keep"`
		Map       map[string]string `json:"map,omitempty" faker:"keep"`
	}

	firstName := "Heino van der Laien"
	m := map[string]string{"foo": "bar"}
	test := TestStruct{
		FirstName: firstName,
		Map:       m,
	}

	err := FakeData(&test)
	if err != nil {
		t.Error("expected not error, but got: ", err)
	}

	if test.FirstName != firstName {
		t.Fatalf("expected: %s, but got: %s", firstName, test.FirstName)
	}

	for k, v := range m {
		if test.Map[k] != v {
			t.Fatalf("expected: %s, but got: %s", m, test.Map)
		}
	}

	if test.Email == "" {
		t.Error("expected filled but got empty")
	}
}

func TestItThrowsAnErrorWhenKeepIsUsedOnIncomparableType(t *testing.T) {
	type TypeStructWithStruct struct {
		Struct struct{} `faker:"first_name_male,keep"`
	}
	type TypeStructWithSlice struct {
		Slice []string `faker:"first_name_male,keep"`
	}
	type TypeStructWithArray struct {
		Array [4]string `faker:"first_name_male,keep"`
	}

	withStruct := TypeStructWithStruct{}
	withSlice := TypeStructWithSlice{}
	withArray := TypeStructWithArray{}

	for _, item := range []interface{}{withArray, withStruct, withSlice} {
		err := FakeData(&item)
		if err == nil {
			t.Errorf("expected error, but got nil")
		}
	}
}

func TestItThrowsAnErrorWhenPointerToInterfaceIsUsed(t *testing.T) {
	type PtrToInterface struct {
		Interface *interface{}
	}

	interfacePtr := PtrToInterface{}

	err := FakeData(&interfacePtr)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func TestItThrowsAnErrorWhenZeroValueWithKeepAndUnsupportedTagIsUsed(t *testing.T) {
	type String struct {
		StringVal string `faker:"keep,unsupported"`
	}

	val := String{}

	err := FakeData(&val)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func TestUnique(t *testing.T) {
	type UniqueStruct struct {
		StringVal string `faker:"word,unique"`
		IntVal    int    `faker:"unique"`
	}

	for i := 0; i < 50; i++ {
		val := UniqueStruct{}
		err := FakeData(&val)
		if err != nil {
			t.Fatal("can't fake the unique data", err)
		}
	}

	found := []interface{}{}
	for _, v := range uniqueValues["word"] {
		for _, f := range found {
			if f == v {
				t.Errorf("expected unique values, found \"%s\" at least twice", v)
				ResetUnique()
				return
			}
		}
		found = append(found, v)
	}

	ResetUnique()
}

func TestUniqueReset(t *testing.T) {
	type String struct {
		StringVal string `faker:"word,unique"`
	}

	for i := 0; i < 20; i++ {
		val := String{}
		err := FakeData(&val)
		if err != nil {
			t.Fatal("can't fake the unique data", err)
		}
	}

	ResetUnique()
	length := len(uniqueValues)
	if length > 0 {
		t.Errorf("expected empty uniqueValues map, but got a length of %d", length)
	}
}

func TestUniqueFailure(t *testing.T) {
	type String struct {
		StringVal string `faker:"word,unique"`
	}

	hasError := false
	length := len(wordList) + 1
	for i := 0; i < length; i++ {
		val := String{}
		err := FakeData(&val)
		if err != nil {
			hasError = true
			break
		}
	}

	ResetUnique()
	if !hasError {
		t.Errorf("expected error, but got nil")
	}
}

func TestOneOfTag__GoodInputs(t *testing.T) {

	type CustomOneString struct {
		PaymentType string `faker:"oneof: credit card"`
	}

	t.Run("creates one of the desired string values", func(t *testing.T) {
		a := CustomOneString{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		if a.PaymentType != "credit card" {
			t.Errorf(
				"expected credit card but got %v",
				a.PaymentType,
			)
		}
	})

	type CustomTwoString struct {
		PaymentType string `faker:"oneof: credit card, paypal"`
	}

	t.Run("creates one of the desired string values", func(t *testing.T) {
		a := CustomTwoString{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		one := a.PaymentType == "credit card"
		two := a.PaymentType == "paypal"

		if !one && !two {
			t.Errorf(
				"expected either %v or %v but got %v",
				"credit card",
				"paypal",
				a.PaymentType,
			)
		}
	})

	type CustomMultiString struct {
		PaymentType string `faker:"oneof: cc, check, paypal, bank account"`
	}
	t.Run("creates only one of the desired string values from many", func(t *testing.T) {
		a := CustomMultiString{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		one := a.PaymentType == "cc"
		two := a.PaymentType == "paypal"
		three := a.PaymentType == "check"
		four := a.PaymentType == "bank account"

		if !one && !two && !three && !four {
			t.Errorf(
				"expected either %v or %v or %v or %v but got %v",
				"cc",
				"paypal",
				"check",
				"bank account",
				a.PaymentType,
			)
		}
	})

	type CustomOneofInt1 struct {
		Age int `faker:"oneof: 16, 18, 21"`
	}

	t.Run("should pick one of the number args", func(t *testing.T) {
		a := CustomOneofInt1{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		one := a.Age == 16
		two := a.Age == 18
		three := a.Age == 21
		actual := a.Age
		if !one && !two && !three {
			t.Errorf(
				"expected either %v, %v, or %v, but got %v",
				16, 18, 21, actual,
			)
		}
	})

	type CustomOneofInt struct {
		Age int `faker:"oneof: 15"`
	}

	t.Run("errors when tag is not used correctly int only one argument", func(t *testing.T) {
		a := CustomOneofInt{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		if a.Age != 15 {
			t.Errorf(
				"expected either %v, but got %v",
				15, a.Age,
			)
		}
	})

	type CustomFloat1 struct {
		Price float32 `faker:"oneof: 3.14, 15.92"`
	}

	t.Run("correctly picks one of the float32s", func(t *testing.T) {
		a := CustomFloat1{}
		err := FakeData(&a)
		if err != nil {
			t.Error("expected no error but got ", err)
		}
		one := a.Price == 3.14
		two := a.Price == 15.92

		if !one && !two {
			t.Errorf("expected either %v or %v but got %v", 3.14, 15.92, a.Price)
		}
	})

	type CustomFloat2 struct {
		Price float32 `faker:"oneof: 1848205.48483727"`
	}

	t.Run("errors when tag is not used correctly float32 only one argument", func(t *testing.T) {
		a := CustomFloat2{}
		err := FakeData(&a)
		if err != nil {
			t.Error("expected no error but got ", err)
		}
		if a.Price != float32(1848205.48483727) {
			t.Errorf("expected either %v but got %v", 1848205.48483727, a.Price)
		}
	})

	type CustomFloat3 struct {
		Price float64 `faker:"oneof: 1848205.48483727"`
	}

	t.Run("errors when tag is not used correctly float64 only one argument", func(t *testing.T) {
		a := CustomFloat3{}
		err := FakeData(&a)
		if err != nil {
			t.Error("expected no error but got ", err)
		}
		if a.Price != float64(1848205.48483727) {
			t.Errorf("expected either %v but got %v", 1848205.48483727, a.Price)
		}
	})

	type CustomFloat6 struct {
		Price float64 `faker:"oneof: 34566872.57446732, 969525372.57563314"`
	}

	t.Run("correctly picks one of the float64s", func(t *testing.T) {
		a := CustomFloat6{}
		err := FakeData(&a)
		if err != nil {
			t.Error("expected no error but got ", err)
		}
		const first = 34566872.57446732
		const second = 969525372.57563314
		one := a.Price == first
		two := a.Price == second

		if !one && !two {
			t.Errorf("expected either %v or %v but got %v", first, second, a.Price)
		}
	})

	type CustomTypeLotsOfInts struct {
		Age1 int64  `faker:"oneof: 1, 2"`
		Age2 int32  `faker:"oneof: 3, 5"`
		Age3 int16  `faker:"oneof: 8, 13"`
		Age4 int8   `faker:"oneof: 21, 34"`
		Age5 int    `faker:"oneof: 55, 89"`
		Age6 uint64 `faker:"oneof: 2, 4"`
		Age7 uint32 `faker:"oneof: 6, 8"`
		Age8 uint16 `faker:"oneof: 10, 12"`
		Age9 uint8  `faker:"oneof: 3, 5"`
		Age0 uint   `faker:"oneof: 7, 9"`
	}

	t.Run("Should support all the int types", func(t *testing.T) {
		a := CustomTypeLotsOfInts{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
	})

	type CustomTypeLotsOfPtrNumbers struct {
		Age1 *int64  `faker:"oneof: 1, 2"`
		Age2 *int32  `faker:"oneof: 3, 5"`
		Age3 *int16  `faker:"oneof: 8, 3"`
		Age4 *int8   `faker:"oneof: 7, 9"`
		Age5 *int    `faker:"oneof: 6, 2"`
		Age6 *uint64 `faker:"oneof: 2, 4"`
		Age7 *uint32 `faker:"oneof: 6, 8"`
		Age8 *uint16 `faker:"oneof: 9, 6"`
		Age9 *uint8  `faker:"oneof: 3, 5"`
		Age0 *uint   `faker:"oneof: 1, 4"`
	}

	t.Run("Should support all the ptr number types", func(t *testing.T) {
		a := CustomTypeLotsOfPtrNumbers{}
		err := FakeData(&a)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		val := reflect.ValueOf(a)

		for i := 0; i < val.Type().NumField(); i++ {
			if val.Field(i).IsZero() {
				t.Errorf("%s: expected non-nil but got %v", val.Type().Field(i).Name, val.Field(i).Interface())
				continue
			}
			strVal := fmt.Sprintf("%d", val.Field(i).Elem().Interface())
			if len(strVal) != 1 {
				t.Errorf("%s: expected [0,9] but got %s", val.Type().Field(i).Name, strVal)
			}
		}

	})
}

func TestOneOfTag__BadInputsForFloats(t *testing.T) {

	type CustomWrongFloat1 struct {
		Price float32 `faker:"oneof:"`
	}

	t.Run("errors when tag is not used correctly no args float32", func(t *testing.T) {
		a := CustomWrongFloat1{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrNotEnoughTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat2 struct {
		Price float32 `faker:"oneof: 15.5: 18.9, 35.4747"`
	}

	t.Run("errors when tag is not used correctly float32 invalid argument separator", func(t *testing.T) {
		a := CustomWrongFloat2{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat3 struct {
		Price float32 `faker:"oneof: 1648.4564673, 894572.997376, oops"`
	}

	t.Run("errors when tag is not used correctly float32 invalid argument type", func(t *testing.T) {
		a := CustomWrongFloat3{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat5 struct {
		Price float32 `faker:"oneof: 15,,16,17"`
	}

	t.Run("errors when tag is not used correctly float32 only one argument", func(t *testing.T) {
		a := CustomWrongFloat5{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrDuplicateSeparator
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat7 struct {
		Price float64 `faker:"oneof:"`
	}

	t.Run("errors when tag is not used correctly no args float64", func(t *testing.T) {
		a := CustomWrongFloat7{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
			return
		}
		actual := err.Error()
		expected := fakerErrors.ErrNotEnoughTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat8 struct {
		Price float64 `faker:"oneof: 157285.842725: 184028.474729, 3574626.4747"`
	}

	t.Run("errors when tag is not used correctly float64 invalid argument separator", func(t *testing.T) {
		a := CustomWrongFloat8{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat9 struct {
		Price float64 `faker:"oneof: 1648.4564673, 894572.997376, oops"`
	}

	t.Run("errors when tag is not used correctly float64 invalid argument type", func(t *testing.T) {
		a := CustomWrongFloat9{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error", err)
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongFloat11 struct {
		Price float64 `faker:"oneof: 15,,16,17"`
	}

	t.Run("errors when tag is not used correctly float64 only one argument", func(t *testing.T) {
		a := CustomWrongFloat11{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrDuplicateSeparator
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

}

func TestOneOfTag__BadInputsForStrings(t *testing.T) {

	type CustomOneofWrongString struct {
		PaymentType string `faker:"oneof:"`
	}

	t.Run("errors when tag is not used correctly string no args", func(t *testing.T) {
		a := CustomOneofWrongString{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrNotEnoughTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomOneofWrongString1 struct {
		PaymentType string `faker:"oneof"`
	}

	t.Run("errors when tag is not used correctly string no args or even colon separator", func(t *testing.T) {
		a := CustomOneofWrongString1{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomOneofWrongString2 struct {
		PaymentType string `faker:"oneof: cc: check, bank"`
	}

	t.Run("errors when tag is not used correctly string invalid argument separator", func(t *testing.T) {
		a := CustomOneofWrongString2{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongString4 struct {
		PaymentType string `faker:"oneof: ,,,cc, credit card,,"`
	}

	t.Run("errors when tag is not used correctly string duplicate separator", func(t *testing.T) {
		a := CustomWrongString4{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrDuplicateSeparator
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

}

func TestOneOfTag__BadInputsForInts(t *testing.T) {

	type CustomTypeInt64Wrong struct {
		Age int64 `faker:"oneof: 1_000_000, oops"`
		Avg int64 `faker:"boundary_start=31, boundary_end=88"`
	}

	t.Run("should error for int64 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeInt64Wrong{}
		err := FakeData(&a)
		t.Log(a.Age, a.Avg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeInt32Wrong struct {
		Age int32 `faker:"oneof: 1_000_000, oops"`
	}

	t.Run("should error for int32 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeInt32Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeInt16Wrong struct {
		Age int16 `faker:"oneof: 1_000, oops"`
	}

	t.Run("should error for int16 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeInt16Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeInt8Wrong struct {
		Age int8 `faker:"oneof: 250, oops"`
	}

	t.Run("should error for int8 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeInt8Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeUint64Wrong struct {
		Age uint64 `faker:"oneof: 250_000_000, oops"`
	}

	t.Run("should error for uint64 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeUint64Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeUint32Wrong struct {
		Age uint32 `faker:"oneof: 2_000_000, oops"`
	}

	t.Run("should error for uint32 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeUint32Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeUint16Wrong struct {
		Age uint16 `faker:"oneof: 2_000, oops"`
	}

	t.Run("should error for uint16 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeUint16Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomTypeUint8Wrong struct {
		Age uint8 `faker:"oneof: 400, oops"`
	}

	t.Run("should error for uint8 with bad tag arguments", func(t *testing.T) {
		a := CustomTypeUint8Wrong{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	type CustomOneofWrongInt struct {
		Age int `faker:"oneof:"`
	}

	t.Run("errors when tag is not used correctly no args int", func(t *testing.T) {
		a := CustomOneofWrongInt{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrNotEnoughTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomOneofWrongInt2 struct {
		Age int `faker:"oneof: 15: 18, 35"`
	}

	t.Run("errors when tag is not used correctly int invalid argument separator", func(t *testing.T) {
		a := CustomOneofWrongInt2{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomOneofWrongInt3 struct {
		Age int `faker:"oneof: 15, 18, oops"`
	}

	t.Run("errors when tag is not used correctly int invalid argument type", func(t *testing.T) {
		a := CustomOneofWrongInt3{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomWrongInt5 struct {
		Age int `faker:"oneof: 15,,16,17"`
	}

	t.Run("errors when tag is not used correctly int only one argument", func(t *testing.T) {
		a := CustomWrongInt5{}
		err := FakeData(&a)
		if err == nil {
			t.Fatal("expected error, but got no error")
		}
		actual := err.Error()
		expected := fakerErrors.ErrDuplicateSeparator
		if actual != expected {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})

	type CustomTypeNegativeUnsigned struct {
		Age uint `faker:"oneof: -45, -42"`
	}

	t.Run("passing a negative int to an unsigned int should cause error", func(t *testing.T) {
		a := CustomTypeNegativeUnsigned{}
		err := FakeData(&a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		expected := fakerErrors.ErrUnsupportedTagArguments
		actual := err.Error()
		if expected != actual {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestFakeData3(t *testing.T) {
	t.Run("test nil pointer", func(t *testing.T) {
		var a *int
		err := FakeData(a)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected1 := "Use sample:=new"
		expected2 := "faker.FakeData(sample) instead"
		if !strings.Contains(actual, expected1) && !strings.Contains(actual, expected2) {
			t.Errorf("expected %v to contain %v & %v", actual, expected1, expected2)
		}
	})
}

// getStringLen for language independent string length
func utfLen(value string) int {
	var r int
	fmt.Println(len(value))
	for i := 0; i < len(value); {
		_, size := utf8.DecodeRuneInString(value[i:])
		i += size
		r++
	}
	return r
}

func TestRandomMaxMinMapSliceSize(t *testing.T) {
	type SliceMap struct {
		Slice []int
		Map   map[string]int
	}

	for _, c := range []struct {
		max, min, expect int
	}{
		{2, 1, 1}, // [1,2) => always 1
		{2, 2, 2},
		{2, 3, 3}, // if min >= max, result will always be min
	} {

		s := SliceMap{}
		err := FakeData(&s, options.WithRandomMapAndSliceMaxSize(uint(c.max)),
			options.WithRandomMapAndSliceMinSize(uint(c.min)))
		if err != nil {
			t.Error(err)
		}

		if len(s.Map) != c.expect {
			t.Errorf("map (len:%d) not expect length with test case %+v\n", len(s.Map), c)
		}

		if len(s.Slice) != c.expect {
			t.Errorf("slice (len:%d) not expect length with test case %+v\n", len(s.Slice), c)
		}
	}
}

func TestRandomMapSliceSize(t *testing.T) {
	// test if old func behaves the same
	type SliceMap struct {
		Slice []int
		Map   map[string]int
	}
	expect := 5
	for i := 0; i < 10; i++ {
		s := SliceMap{}
		err := FakeData(&s, options.WithRandomMapAndSliceMaxSize(uint(expect)))
		if err != nil {
			t.Error(err)
		}

		if len(s.Map) >= 5 {
			t.Errorf("map (len:%d) is greater than expected length %d", len(s.Map), expect)
		}

		if len(s.Slice) >= expect {
			t.Errorf("slice (len:%d) not expect length with test case %+v\n", len(s.Slice), expect)
		}
	}
}

func TestWithFieldsToIgnore(t *testing.T) {
	a := AStruct{}
	if err := FakeData(&a, options.WithFieldsToIgnore("Height", "Name")); err != nil {
		t.Error(err)
	}

	if a.Height != 0 {
		t.Errorf("Height expected to be ignored")
	}
	if a.AnotherStruct.Name != "" {
		t.Errorf("Name expected to be ignored")
	}
	if a.AnotherStruct.Image == "" {
		t.Errorf("other fields are affected")
	}
}

func TestWithFieldProvider(t *testing.T) {
	a := AStruct{}
	const heightVal = int64(123)
	const nameVal = "some string"
	if err := FakeData(&a,
		options.WithCustomFieldProvider("Height", func() (interface{}, error) {
			return heightVal, nil
		}),
		options.WithCustomFieldProvider("Name", func() (interface{}, error) {
			return nameVal, nil
		}),
	); err != nil {
		t.Error(err)
	}

	if a.Height != heightVal {
		t.Errorf("expected Height %d, got %d", heightVal, a.Height)
	}
	if a.AnotherStruct.Name != nameVal {
		t.Errorf("expected Name %q, got %q", nameVal, a.AnotherStruct.Name)
	}

	// when provider fails
	if err := FakeData(&a, options.WithCustomFieldProvider("Height", func() (interface{}, error) {
		return nil, fmt.Errorf("test")
	})); err == nil {
		t.Errorf("expected an error, but got nil")
	}
}

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type GeneralTreeNode struct {
	Val      int
	Children []*GeneralTreeNode
}

func TestFakeData_RecursiveType(t *testing.T) {
	flatTree := func(root *BinaryTreeNode) []*BinaryTreeNode {
		if root == nil {
			return nil
		}
		ans := []*BinaryTreeNode{root}
		for i := 0; i < len(ans); i++ {
			if ans[i].Left != nil {
				ans = append(ans, ans[i].Left)
			}
			if ans[i].Right != nil {
				ans = append(ans, ans[i].Right)
			}
		}
		return ans
	}
	// depth 1
	var root *BinaryTreeNode
	if err := FakeData(&root); err != nil {
		t.Errorf("%+v", err)
		t.FailNow()
	}
	nodes := flatTree(root)
	if len(nodes) != 3 {
		t.Errorf("expect 3 node, got %d", len(nodes))
		t.FailNow()
	}
	if root == nil || root.Left == nil || root.Left.Left != nil {
		t.Errorf("expect depth: 1")
		t.FailNow()
	}
	// depth 0
	root = nil
	if err := FakeData(&root, options.WithRecursionMaxDepth(0)); err != nil {
		t.Errorf("%+v", err)
		t.FailNow()
	}
	if root == nil || root.Left != nil || root.Right != nil {
		t.Errorf("expect depth: 0")
		t.FailNow()
	}

	// depth 0
	var root2 *GeneralTreeNode
	if err := FakeData(&root2, options.WithRecursionMaxDepth(0)); err != nil {
		t.Errorf("%+v", err)
		t.FailNow()
	}
	if root2 == nil {
		t.Errorf("empty root")
		t.FailNow()
	}
	for _, child := range root2.Children {
		if child != nil {
			t.Errorf("expect depth: 0")
			t.FailNow()
		}
	}
}
