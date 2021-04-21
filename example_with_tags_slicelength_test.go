package faker_test

import (
	"fmt"

	"github.com/cloudquery/faker/v3"
)

// You can set the size for your random slices.
func Example_withTagsSliceLength() {
	// SomeStruct ...
	type SomeStruct struct {
		EmptyList       []string `faker:"slice_len=0"`
		FixedStringList []string `faker:"slice_len=2"`
		FixedIntList    []int64  `faker:"slice_len=4"`
		RandomIntList    []int64
	}

	_ = faker.SetRandomMapAndSliceSize(20) // If no slice_len is set, this sets the max of the random size
	a := SomeStruct{}
	_ = faker.FakeData(&a)
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
