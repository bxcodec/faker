package faker_test

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
)

// Single fake function can be used for retrieving particular values.
func Example_singleFakeData() {

	// Address
	fmt.Println(faker.Latitude())  // => 81.12195
	fmt.Println(faker.Longitude()) // => -84.38158

	// Datetime
	fmt.Println(faker.UnixTime())   // => 1197930901
	fmt.Println(faker.Date())       // => 1982-02-27
	fmt.Println(faker.TimeString()) // => 03:10:25
	fmt.Println(faker.MonthName())  // => February
	fmt.Println(faker.YearString()) // => 1994
	fmt.Println(faker.DayOfWeek())  // => Sunday
	fmt.Println(faker.DayOfMonth()) // => 20
	fmt.Println(faker.Timestamp())  // => 1973-06-21 14:50:46
	fmt.Println(faker.Century())    // => IV
	fmt.Println(faker.Timezone())   // => Asia/Jakarta
	fmt.Println(faker.Timeperiod()) // => PM

	// Internet
	fmt.Println(faker.Email())      // => mJBJtbv@OSAaT.com
	fmt.Println(faker.MacAddress()) // => cd:65:e1:d4:76:c6
	fmt.Println(faker.DomainName()) // => FWZcaRE.org
	fmt.Println(faker.URL())        // => https://www.oEuqqAY.org/QgqfOhd
	fmt.Println(faker.Username())   // => lVxELHS
	fmt.Println(faker.IPv4())       // => 99.23.42.63
	fmt.Println(faker.IPv6())       // => 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
	fmt.Println(faker.Password())   // => dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTP

	// Words and Sentences
	fmt.Println(faker.Word())      // => nesciunt
	fmt.Println(faker.Sentence())  // => Consequatur perferendis voluptatem accusantium.
	fmt.Println(faker.Paragraph()) // => Aut consequatur sit perferendis accusantium voluptatem. Accusantium perferendis consequatur voluptatem sit aut. Aut sit accusantium consequatur voluptatem perferendis. Perferendis voluptatem aut accusantium consequatur sit.

	// Payment
	fmt.Println(faker.CCType())             // => American Express
	fmt.Println(faker.CCNumber())           // => 373641309057568
	fmt.Println(faker.Currency())           // => USD
	fmt.Println(faker.AmountWithCurrency()) // => USD 49257.100

	// Person
	fmt.Println(faker.TitleMale())       // => Mr.
	fmt.Println(faker.TitleFemale())     // => Mrs.
	fmt.Println(faker.FirstName())       // => Whitney
	fmt.Println(faker.FirstNameMale())   // => Kenny
	fmt.Println(faker.FirstNameFemale()) // => Jana
	fmt.Println(faker.LastName())        // => Rohan
	fmt.Println(faker.Name())            // => Mrs. Casandra Kiehn

	// Phone
	fmt.Println(faker.Phonenumber())         // -> 201-886-0269
	fmt.Println(faker.TollFreePhoneNumber()) // => (777) 831-964572
	fmt.Println(faker.E164PhoneNumber())     // => +724891571063

	//  UUID
	fmt.Println(faker.UUIDHyphenated()) // => 8f8e4463-9560-4a38-9b0c-ef24481e4e27
	fmt.Println(faker.UUIDDigit())      // => 90ea6479fd0e4940af741f0a87596b73

	fmt.Println(faker.Word())

	faker.ResetUnique() // Forget all generated unique values
}
