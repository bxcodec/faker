package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type SomeStruct struct {
	Inta             int
	Int8             int8
	Int16            int16
	Int32            int32
	Int64            int64
	Float32          float32
	Float64          float64
	Latitude         float32 `faker:"lat"`
	LATITUDE         float64 `faker:"lat"`
	Long             float32 `faker:"long"`
	LONG             float64 `faker:"long"`
	String           string
	CreditCardType   string `faker:"cc_type"`
	CreditCardNumber string `faker:"cc_number"`
	Email            string `faker:"email"`
	IPV4             string `faker:"ipv4"`
	IPV6             string `faker:"ipv6"`
	Bool             bool
	SString          []string
	SInt             []int
	SInt8            []int8
	SInt16           []int16
	SInt32           []int32
	SInt64           []int64
	SFloat32         []float32
	SFloat64         []float64
	SBool            []bool
	Struct           AStruct
	Time             time.Time
	Stime            []time.Time
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
	Latitude         float32 `faker:"lat"`
	Long             float32 `faker:"long"`
	CreditCardNumber string  `faker:"cc_number"`
	CreditCardType   string  `faker:"cc_type"`
	Email            string  `faker:"email"`
	IPV4             string  `faker:"ipv4"`
	IPV6             string  `faker:"ipv6"`
	PhoneNumber      string  `faker:"phone_number"`
	MacAddress       string  `faker:"mac_address"`
	Url              string  `faker:"url"`
	UserName         string  `faker:"username"`
	ToolFreeNumber   string  `faker:"tool_free_number"`
	E164PhoneNumber  string  `faker:"e_164_phone_number"`
	TitleMale        string  `faker:"title_male"`
	TitleFemale      string  `faker:"title_female"`
	FirstNameMale    string  `faker:"first_name_male"`
	FirstNameFemale  string  `faker:"first_name_female"`
	LastName         string  `faker:"last_name"`
	Name         string  `faker:"name"`
}

func (t TaggedStruct) String() string {
	return fmt.Sprintf(`{
	Latitude: %f,
	Long: %f,
	CreditCardNumber: %s,
	CreditCardType: %s,
	Email: %s,
	IPV4: %s,
	IPV6: %s,
	PhoneNumber: %s,
	MacAddress: %s,
	Url: %s,
	UserName: %s,
	ToolFreeNumber: %s,
	E164PhoneNumber: %s,
	TitleMale: %s,
	TitleFemale: %s,
	FirstNameMale: %s,
	FirstNameFemale: %s,
	LastName: %s,
	Name: %s,
}`, t.Latitude, t.Long, t.CreditCardNumber,
		t.CreditCardType, t.Email, t.IPV4,
		t.IPV6, t.PhoneNumber, t.MacAddress,
		t.Url, t.UserName, t.ToolFreeNumber,
		t.E164PhoneNumber, t.TitleMale, t.TitleFemale,
		t.FirstNameMale, t.FirstNameFemale, t.LastName,
		t.Name)
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
	if err == nil {
		fmt.Println("SomeStruct:")
		fmt.Printf("%+v\n", a)
	}
	var b TaggedStruct
	err = FakeData(&b)
	if err == nil {
		fmt.Println("TaggedStruct:")
		fmt.Printf("%+v\n", b)

	} else {
		fmt.Println(" ER ", err)
	}

	// Example Result :
	// {Int:8906957488773767119 Int8:6 Int16:14 Int32:391219825 Int64:2374447092794071106 String:poraKzAxVbWVkMkpcZCcWlYMd Bool:false SString:[MehdV aVotHsi] SInt:[528955241289647236 7620047312653801973 2774096449863851732] SInt8:[122 -92 -92] SInt16:[15679 -19444 -30246] SInt32:[1146660378 946021799 852909987] SInt64:[6079203475736033758 6913211867841842836 3269201978513619428] SFloat32:[0.019562425 0.12729558 0.36450312] SFloat64:[0.7825838989890364 0.9732903338838912 0.8316541489234004] SBool:[true false true] Struct:{Number:7693944638490551161 Height:6513508020379591917}}

}

func TestSetSliceDataNotFoundType(t *testing.T) {
	if "Slice of string Not Supported Yet" != setSliceData(reflect.ValueOf("")).Error() {
		t.Error("Expected error from func setSliceData")
	}
}

func TestSetDataIfArgumentNotPtr(t *testing.T) {
	temp := struct{}{}
	if "Not a pointer value" != setData(reflect.ValueOf(temp)).Error() {
		t.Error("Expected in arguments not ptr")
	}
}

func TestSetDataIfArgumentPtr(t *testing.T) {
	temp := &struct{}{}
	if "Unsupported kind: ptr Change Without using * (pointer) in Field of *struct {}" != setData(reflect.ValueOf(&temp)).Error() {
		t.Error("Exptected error Unsupported kind ptr")
	}
}

func TestSetDataIfArgumentNotHaveReflect(t *testing.T) {
	temp := func() {}
	if "Unsupported kind: func" != setData(reflect.ValueOf(&temp)).Error() {
		t.Error("Exptected error Unsupported kind")
	}
}

func TestSetDataErrorDataParseTag(t *testing.T) {
	temp := &struct {
		test string `faker:"test"`
	}{}
	if "String Tag unsupported" != setData(reflect.ValueOf(temp)).Error() {
		t.Error("Exptected error Unsupported tag")
	}
}

func TestSetDataWithTagIfFirstArgumentNotPtr(t *testing.T) {
	temp := struct{}{}
	if "Not a pointer value" != setDataWithTag(reflect.ValueOf(temp), "").Error() {
		t.Error("Expected in arguments not ptr")
	}
}

func TestSetDataWithTagIfFirstArgumentSlice(t *testing.T) {
	temp := []int{}
	if setDataWithTag(reflect.ValueOf(&temp), "") != nil {
		t.Error("Not expected errors if first argument slice type")
	}
}

func TestSetDataWithTagIfFirstArgumentNotFound(t *testing.T) {
	temp := struct{}{}
	if setDataWithTag(reflect.ValueOf(&temp), "") != nil {
		t.Error("First argument is struct type, expected return nil")
	}
}

func TestUserDefinedFloatNotFoundTag(t *testing.T) {
	temp := struct{}{}

	if userDefinedFloat(reflect.ValueOf(&temp), "") == nil {
		t.Error("Not expected errors")
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
