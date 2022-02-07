package faker_test

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// You can set length for your random strings also set boundary for your integers.
func Example_withTagsLang() {
	// SomeStruct ...
	type SomeStruct struct {
		StringENG string `faker:"lang=eng"`
		StringCHI string `faker:"lang=chi"`
		StringRUS string `faker:"lang=rus"`
		StringJPN string `faker:"lang=jpn"`
		StringKOR string `faker:"lang=kor"`
		StringEMJ string `faker:"lang=emj"`
	}

	a := SomeStruct{}
	_ = faker.SetRandomStringLength(5)
	_ = faker.FakeData(&a)
	fmt.Printf("%+v", a)
	// Result:
	/*
		   {
			   StringENG:VVcaPS
			   StringCHI: 随机字符串
			   StringRUS:ваЩфз
			   StringJPN:びゃほぱヒてふ
			   StringKOR:텻밚쨋큊몉
			   StringEMJ:🐅😄🕢🍪🐡
		   }
	*/
}
