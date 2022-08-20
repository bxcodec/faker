package faker_test

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
)

// SomeStructWithUnique ...
type SomeStructWithUnique struct {
	Word string `faker:"word,unique"`
}

func Example_withTagsAndUnique() {
	for i := 0; i < 5; i++ { // Generate 5 structs having a unique word
		a := SomeStructWithUnique{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", a)
	}
	faker.ResetUnique() // Forget all generated unique values. Allows to start generating another unrelated dataset.

	// Result:
	//	{Word:nobis}{Word:recusandae}{Word:praesentium}{Word:doloremque}{Word:non}
}
