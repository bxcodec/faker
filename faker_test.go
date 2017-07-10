package faker

import (
	"fmt"
	"testing"
	"time"
)

type SomeStruct struct {
	// Inta     int
	// Int8     int8
	// Int16    int16
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

func TestFakerData(t *testing.T) {
	var a SomeStruct
	err := FakeData(&a)
	if err == nil {
		fmt.Printf("%+v", a)
	}

	// Example Result :
	// {Int:8906957488773767119 Int8:6 Int16:14 Int32:391219825 Int64:2374447092794071106 String:poraKzAxVbWVkMkpcZCcWlYMd Bool:false SString:[MehdV aVotHsi] SInt:[528955241289647236 7620047312653801973 2774096449863851732] SInt8:[122 -92 -92] SInt16:[15679 -19444 -30246] SInt32:[1146660378 946021799 852909987] SInt64:[6079203475736033758 6913211867841842836 3269201978513619428] SFloat32:[0.019562425 0.12729558 0.36450312] SFloat64:[0.7825838989890364 0.9732903338838912 0.8316541489234004] SBool:[true false true] Struct:{Number:7693944638490551161 Height:6513508020379591917}}

}

func BenchmarkFakerData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := SomeStruct{}
		err := FakeData(&a)
		if err != nil {
			b.Fatal(err)
		}
	}
}
