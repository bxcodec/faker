package faker_test

import (
	"github.com/bxcodec/faker/v4"
	"github.com/bxcodec/faker/v4/options"
)

// Single fake function can be used for retrieving particular values.
func Example_singleFakeData() {

	// Address
	faker.Latitude(options.DefaultOption())  // => 81.12195
	faker.Longitude(options.DefaultOption()) // => -84.38158

	// Datetime
	faker.UnixTime(options.DefaultOption())   // => 1197930901
	faker.Date(options.DefaultOption())       // => 1982-02-27
	faker.TimeString(options.DefaultOption()) // => 03:10:25
	faker.MonthName(options.DefaultOption())  // => February
	faker.YearString(options.DefaultOption()) // => 1994
	faker.DayOfWeek(options.DefaultOption())  // => Sunday
	faker.DayOfMonth(options.DefaultOption()) // => 20
	faker.Timestamp(options.DefaultOption())  // => 1973-06-21 14:50:46
	faker.Century(options.DefaultOption())    // => IV
	faker.Timezone(options.DefaultOption())   // => Asia/Jakarta
	faker.Timeperiod(options.DefaultOption()) // => PM

	// Internet
	faker.Email(options.DefaultOption())      // => mJBJtbv@OSAaT.com
	faker.MacAddress(options.DefaultOption()) // => cd:65:e1:d4:76:c6
	faker.DomainName(options.DefaultOption()) // => FWZcaRE.org
	faker.URL(options.DefaultOption())        // => https://www.oEuqqAY.org/QgqfOhd
	faker.Username(options.DefaultOption())   // => lVxELHS
	faker.IPv4(options.DefaultOption())       // => 99.23.42.63
	faker.IPv6(options.DefaultOption())       // => 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
	faker.Password(options.DefaultOption())   // => dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTP

	// Words and Sentences
	faker.Word(options.DefaultOption())      // => nesciunt
	faker.Sentence(options.DefaultOption())  // => Consequatur perferendis voluptatem accusantium.
	faker.Paragraph(options.DefaultOption()) // => Aut consequatur sit perferendis accusantium voluptatem. Accusantium perferendis consequatur voluptatem sit aut. Aut sit accusantium consequatur voluptatem perferendis. Perferendis voluptatem aut accusantium consequatur sit.

	// Payment
	faker.CCType(options.DefaultOption())             // => American Express
	faker.CCNumber(options.DefaultOption())           // => 373641309057568
	faker.Currency(options.DefaultOption())           // => USD
	faker.AmountWithCurrency(options.DefaultOption()) // => USD 49257.100

	// Person
	faker.TitleMale(options.DefaultOption())       // => Mr.
	faker.TitleFemale(options.DefaultOption())     // => Mrs.
	faker.FirstName(options.DefaultOption())       // => Whitney
	faker.FirstNameMale(options.DefaultOption())   // => Kenny
	faker.FirstNameFemale(options.DefaultOption()) // => Jana
	faker.LastName(options.DefaultOption())        // => Rohan
	faker.Name(options.DefaultOption())            // => Mrs. Casandra Kiehn

	// Phone
	faker.Phonenumber(options.DefaultOption())         // -> 201-886-0269
	faker.TollFreePhoneNumber(options.DefaultOption()) // => (777) 831-964572
	faker.E164PhoneNumber(options.DefaultOption())     // => +724891571063

	//  UUID
	faker.UUIDHyphenated(options.DefaultOption()) // => 8f8e4463-9560-4a38-9b0c-ef24481e4e27
	faker.UUIDDigit(options.DefaultOption())      // => 90ea6479fd0e4940af741f0a87596b73

	faker.Word(options.DefaultOption())

	faker.ResetUnique() // Forget all generated unique values
}
