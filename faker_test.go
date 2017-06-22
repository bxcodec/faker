package faker

import (
	"fmt"
	"testing"
)

func TestFakerData(t *testing.T) {
	// a := SomeStruct{}
	var a SomeStruct
	err := FakeData(&a)

	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
	}
}
