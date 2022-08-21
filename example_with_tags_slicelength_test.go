package faker_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/bxcodec/faker/v4/pkg/options"
)

// SomeStructWithSliceLength ...
type SomeStructWithSliceLength struct {
	EmptyList       []string `faker:"slice_len=0"`
	FixedStringList []string `faker:"slice_len=2"`
	FixedIntList    []int64  `faker:"slice_len=4"`
	RandomIntList   []int64
}

// You can set the size for your random slices.
func Example_withTagsSliceLength() {
	a := SomeStructWithSliceLength{}
	_ = faker.FakeData(&a, options.WithRandomMapAndSliceMaxSize(20)) // If no slice_len is set, this sets the max of the random size
	fmt.Printf("%+v", a)
	// Result:
	/*
		   {
		       EmptyList:[]
		       FixedStringList:[
			           geHYIpEoQhQdijFooVEAOyvtTwJOofbQPJdbHvEEdjueZaKIgI
			           WVJBBtmrrVccyIydAiLSkMwWbFzFMEotEXsyUXqcmBTVORlkJK
			   ]
			   FixedIntList:[10,25,60,15]
		       RandomIntList:[5,16,134,6235,123,53,123] //Random len() with a max of 20
		   }
	*/
}

// You can set the size for your random slices.
func Test_withTagsSliceLength(t *testing.T) {
	a := SomeStructWithSliceLength{}
	err := faker.FakeData(&a, options.WithRandomMapAndSliceMaxSize(20)) // If no slice_len is set, this sets the max of the random size
	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}
	if len(a.EmptyList) != 0 {
		t.Errorf("want %v, got %v", 0, len(a.EmptyList))
	}
	if len(a.FixedStringList) > 2 {
		t.Errorf("want <=%v, got %v", 2, len(a.FixedStringList))
	}
	if len(a.FixedIntList) > 4 {
		t.Errorf("want <=%v, got %v", 4, len(a.FixedIntList))
	}
}
