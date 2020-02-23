package faker_test

import "github.com/bxcodec/faker/v3"

// Single fake function can be used for retrieving particular values.
func Example_singleFakeData() {

	// Address
	faker.Latitude()  // => 81.12195
	faker.Longitude() // => -84.38158

	// Datetime
	faker.UnixTime()   // => 1197930901
	faker.Date()       // => 1982-02-27
	faker.TimeString() // => 03:10:25
	faker.MonthName()  // => February
	faker.YearString() // => 1994
	faker.DayOfWeek()  // => Sunday
	faker.DayOfMonth() // => 20
	faker.Timestamp()  // => 1973-06-21 14:50:46
	faker.Century()    // => IV
	faker.Timezone()   // => Asia/Jakarta
	faker.Timeperiod() // => PM

	// Internet
	faker.Email()      // => mJBJtbv@OSAaT.com
	faker.MacAddress() // => cd:65:e1:d4:76:c6
	faker.DomainName() // => FWZcaRE.org
	faker.URL()        // => https://www.oEuqqAY.org/QgqfOhd
	faker.Username()   // => lVxELHS
	faker.IPv4()       // => 99.23.42.63
	faker.IPv6()       // => 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
	faker.Password()   // => dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTP

	// Words and Sentences
	faker.Word()      // => nesciunt
	faker.Sentence()  // => Consequatur perferendis voluptatem accusantium.
	faker.Paragraph() // => Aut consequatur sit perferendis accusantium voluptatem. Accusantium perferendis consequatur voluptatem sit aut. Aut sit accusantium consequatur voluptatem perferendis. Perferendis voluptatem aut accusantium consequatur sit.

	// Payment
	faker.CCType()             // => American Express
	faker.CCNumber()           // => 373641309057568
	faker.Currency()           // => USD
	faker.AmountWithCurrency() // => USD 49257.100

	// Person
	faker.TitleMale()       // => Mr.
	faker.TitleFemale()     // => Mrs.
	faker.FirstName()       // => Whitney
	faker.FirstNameMale()   // => Kenny
	faker.FirstNameFemale() // => Jana
	faker.LastName()        // => Rohan
	faker.Name()            // => Mrs. Casandra Kiehn

	// Phone
	faker.Phonenumber()         // -> 201-886-0269
	faker.TollFreePhoneNumber() // => (777) 831-964572
	faker.E164PhoneNumber()     // => +724891571063

	//  UUID
	faker.UUIDHyphenated() // => 8f8e4463-9560-4a38-9b0c-ef24481e4e27
	faker.UUIDDigit()      // => 90ea6479fd0e4940af741f0a87596b73

	// Unique values
	faker.SetGenerateUniqueValues(true) // Enable unique data generation on single fake data functions
	faker.Word()
	// ...
	faker.SetGenerateUniqueValues(false) // Disable unique data generation on single fake data functions
	faker.ResetUnique()                  // Forget all generated unique values

}
