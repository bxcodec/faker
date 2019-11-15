package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

const (
	someStructLen           = 2
	someStructBoundaryStart = 5
	someStructBoundaryEnd   = 10
)

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

	ASString []string          `faker:"len=2"`
	SString  string            `faker:"len=2"`
	MSString map[string]string `faker:"len=2"`
	MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
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
		t.IPV6, t.Password, t.PhoneNumber, t.MacAddress,
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
	fmt.Printf("%+v\n", a)

	var b TaggedStruct
	err = FakeData(&b)

	if err != nil {
		t.Error("Expected NoError, but Got Err: ", err)
	}

	fmt.Println("TaggedStruct:")
	fmt.Printf("%+v\n", b)

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
		Test string `faker:"test"`
	}{}
	fmt.Printf("%+v ", temp)
	if err := FakeData(temp); err == nil {
		t.Error("Exptected error Unsupported tag, but got nil")
	}
}

func TestSetDataErrorDataParseTagIntType(t *testing.T) {
	temp := &struct {
		Test int `faker:"test"`
	}{}

	if err := FakeData(temp); err == nil {
		t.Error("Expected error Unsupported tag, but got nil")
	}
}

func TestSetRandomStringLength(t *testing.T) {
	someStruct := SomeStruct{}
	if err := SetRandomStringLength(-1); err == nil {
		t.Error("Random string len must not accept lower than 0 as a size")
	}
	strLen := 5
	if err := SetRandomStringLength(strLen); err != nil {
		t.Error("SetRandomStringLength method is corrupted.")
	}
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}
	if len(someStruct.StringValue) > strLen {
		t.Error("SetRandomStringLength did not work.")
	}
}

func TestSetRandomNumberBoundaries(t *testing.T) {
	someStruct := SomeStruct{}
	if err := SetRandomNumberBoundaries(10, 0); err == nil {
		t.Error("Start must be smaller than end value")
	}
	boundary := numberBoundary{start: 10, end: 90}
	if err := SetRandomNumberBoundaries(boundary.start, boundary.end); err != nil {
		t.Error("SetRandomNumberBoundaries method is corrupted.")
	}
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}
	if someStruct.Inta >= boundary.end || someStruct.Inta < boundary.start {
		t.Errorf("%d must be between [%d,%d)", someStruct.Inta, boundary.start, boundary.end)
	}
}

func TestSetRandomMapAndSliceSize(t *testing.T) {
	someStruct := SomeStruct{}
	if err := SetRandomMapAndSliceSize(-1); err == nil {
		t.Error("Random Map and Slice must not accept lower than 0 as a size")
	}
	size := 5
	if err := SetRandomMapAndSliceSize(size); err != nil {
		t.Error("SetRandomMapAndSliceSize method is corrupted.")
	}
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}
	if len(someStruct.MapStringStruct) > size || len(someStruct.SBool) > size {
		t.Error("SetRandomMapAndSliceSize did not work.")
	}
}

func TestSetNilIfLenIsZero(t *testing.T) {
	someStruct := SomeStruct{}
	SetNilIfLenIsZero(true)
	testRandZero = true
	if err := FakeData(&someStruct); err != nil {
		t.Error("Fake data generation has failed")
	}
	if someStruct.MapStringString != nil && someStruct.MapStringStruct != nil &&
		someStruct.MapStringStructPointer != nil {
		t.Error("Map has to be nil")
	}
	if someStruct.Stime != nil && someStruct.SBool != nil {
		t.Error("Array has to be nil")
	}
	testRandZero = false
}

func TestBoundaryAndLen(t *testing.T) {
	iterate := 10
	someStruct := SomeStructWithLen{}
	for i := 0; i < iterate; i++ {
		if err := FakeData(&someStruct); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.Int8)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.Int16)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.Int32)); err != nil {
			t.Error(err)
		}
		if err := validateRange(someStruct.Inta); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.Int64)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.UInt8)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.UInt16)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.UInt32)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.UInta)); err != nil {
			t.Error(err)
		}
		if err := validateRange(int(someStruct.UInt64)); err != nil {
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
			if err := validateRange(k); err != nil {
				t.Error(err)
			}
			if err := validateRange(v); err != nil {
				t.Error(err)
			}
		}
	}
}

func TestExtractNumberFromTagFail(t *testing.T) {
	notSupportedTypeStruct := &struct {
		Test float32 `faker:"boundary_start=5, boundary_end=10"`
	}{}
	if err := FakeData(&notSupportedTypeStruct); err == nil {
		t.Error(err)
	}
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

func validateRange(value int) error {
	if value < someStructBoundaryStart || value > someStructBoundaryEnd {
		return fmt.Errorf("%d must be between %d and %d", value, someStructBoundaryStart,
			someStructBoundaryEnd)
	}
	return nil
}

func TestSetDataWithTagIfFirstArgumentNotPtr(t *testing.T) {
	temp := struct{}{}
	if setDataWithTag(reflect.ValueOf(temp), "").Error() != "Not a pointer value" {
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

func TestRandomIntOnlyError(t *testing.T) {
	arguments := []int{1, 3, 4, 5, 6}
	_, err := RandomInt(arguments...)
	if err == nil && err.Error() == fmt.Errorf(ErrMoreArguments, len(arguments)).Error() {
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
	fmt.Printf(" A value: %+v , Somestruct Value: %+v  ", a, a)

	tagged := new(PointerC)
	err = FakeData(tagged)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	fmt.Printf(" tagged value: %+v , TaggedStruct Value: %+v  ", a, a.PointA.SomeStruct)
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
	fmt.Printf(" A value: %+v , Somestruct Value: %+v  ", a, a)

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
	}
	fmt.Printf(" A value: %+v , SampleStruct Value: %+v  ", a, a)
}

func TestPointerToCustomScalar(t *testing.T) {
	// This test is to ensure that the faker won't panic if trying to fake data on struct that has field
	a := new(CustomInt)
	err := FakeData(a)

	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
	fmt.Printf(" A value: %+v , Custom scalar Value: %+v  ", a, a)
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
	fmt.Printf(" A value: %+v , PointerCustomIntStruct scalar Value: %+v  ", a, a)
}

func TestSkipField(t *testing.T) {
	// This test is to ensure that the faker won't fill field with tag skip

	a := struct {
		ID              int
		ShouldBeSkipped int `faker:"-"`
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

}

func TestTagAlreadyExists(t *testing.T) {
	// This test is to ensure that existing tag cannot be rewritten

	err := AddProvider(EmailTag, func(v reflect.Value) (interface{}, error) {
		return nil, nil
	})

	if err == nil || err.Error() != ErrTagAlreadyExists {
		t.Error("Expected ErrTagAlreadyExists Error,  But Got: ", err)
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

	//Assert
	if sample.FirstName == nil || *sample.FirstName == "" {
		t.Error("Expected filled but got emtpy")
	}
	if sample.Email == nil || *sample.Email == "" {
		t.Error("Expected filled but got emtpy")
	}
	if sample.Latitude == nil || *sample.Latitude == 0 {
		t.Error("Expected filled but got emtpy")
	}
	if sample.Latitude32 == nil || *sample.Latitude32 == 0 {
		t.Error("Expected filled but got emtpy")
	}

	if sample.UnixTime == nil || *sample.UnixTime == 0 {
		t.Error("Expected filled but got emtpy")
	}

	if sample.School == nil || sample.School.Location == "" {
		t.Error("Expected filled but got emtpy")
	}
}

func TestItOverwritesDefaultValueIfKeepIsSet(t *testing.T) {
	type TestStruct struct {
		Email     string `json:"email,omitempty" faker:"email,keep"`
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
