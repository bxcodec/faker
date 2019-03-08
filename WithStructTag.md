## With Tag

Supported tag:

**Internet :**

* Email
* Mac address
* Domain name
* URL
* UserName
* IP Address (IPv4, IPv6 )
* Password


**Payment :**
* Credit Card Type (VISA, MASTERCARD, AMERICAN EXPRESS, DISCOVER, JCB, DINERS CLUB)
* Credit Card Number

**Address :**
* Latitude and Longitude

**Phone :**
* Phone number
* Toll free phone number
* E164PhoneNumber

**Person :**
* Title male
* Title female
* FirstName
* FirstName male
* FirstName female
* LastName
* Name

**DateTime :**
* UnixTime
* Date
* Time
* MonthName
* Year
* DayOfWeek
* DayOfMonth
* Timestamp
* Century
* TimeZone
* TimePeriod

**Lorem :**
* Word
* Sentence
* Paragraph

**Price :**
* Currency
* Amount
* Amount with Currency

**UUID :**
* UUID Digit (32 bytes)
* UUID Hyphenated (36 bytes)

**Skip :**
* \-

```go

package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

// SomeStruct ...
type SomeStruct struct {
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName		   string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated	   string  `faker:"uuid_hyphenated"`
	UUID	           string  `faker:"uuid_digit"`
	Skip		   string  `faker:"-"`
}

func main() {

	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
	/*
		Output  :
		{
			Latitude: 81.12195
			Longitude: -84.38158
			CreditCardType: American Express
			CreditCardNumber: 373641309057568
			Email: mJBJtbv@OSAaT.ru
			DomainName: FWZcaRE.ru,
			IPV4: 99.23.42.63
			IPV6: 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
			Password: dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTPOmNyMFb
			PhoneNumber: 792-153-4861
			MacAddress: cd:65:e1:d4:76:c6
			URL: https://www.oEuqqAY.org/QgqfOhd
			UserName: lVxELHS
			TollFreeNumber: (777) 831-964572
			E164PhoneNumber: +724891571063
			TitleMale: Mr.
			TitleFemale: Queen
			FirstName: Whitney
			FirstNameMale: Kenny
			FirstNameFemale: Jana
			LastName: Rohan
			Name: Miss Casandra Kiehn
			UnixTime: 1197930901
			Date: 1982-02-27
			Time: 03:10:25
			MonthName: February
			Year: 1996
			DayOfWeek: Sunday
			DayOfMonth: 20
			Timestamp: 1973-06-21 14:50:46
			Century: IV
			TimeZone: Canada/Eastern
			TimePeriod: AM
			Word: nesciunt
			Sentence: Consequatur perferendis aut sit voluptatem accusantium.
			Paragraph: Aut consequatur sit perferendis accusantium voluptatem. Accusantium perferendis consequatur voluptatem sit aut. Aut sit accusantium consequatur voluptatem perferendis. Perferendis voluptatem aut accusantium consequatur sit.
			Currency: IRR,
			Amount: 88.990000,
			AmountWithCurrency: XBB 49257.100000,
			UUIDHypenated: 8f8e4463-9560-4a38-9b0c-ef24481e4e27,
			UUID: 90ea6479fd0e4940af741f0a87596b73,
			Skip:
		}
	*/
}

```

## Length And Boundary

---

You can set length for your random strings also set boundary for your integers.
```go
package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

type SomeStruct struct {
	Inta  int   `faker:"boundary_start=5, boundary_end=10"`
	Int8  int8  `faker:"boundary_start=100, boundary_end=1000"`
	Int16 int16 `faker:"boundary_start=123, boundary_end=1123"`
	Int32 int32 `faker:"boundary_start=-10, boundary_end=8123"`
	Int64 int64 `faker:"boundary_start=31, boundary_end=88"`

	UInta  uint   `faker:"boundary_start=35, boundary_end=152"`
	UInt8  uint8  `faker:"boundary_start=5, boundary_end=1425"`
	UInt16 uint16 `faker:"boundary_start=245, boundary_end=2125"`
	UInt32 uint32 `faker:"boundary_start=0, boundary_end=40"`
	UInt64 uint64 `faker:"boundary_start=14, boundary_end=50"`

	ASString []string          `faker:"len=50"`
	SString  string            `faker:"len=25"`
	MSString map[string]string `faker:"len=30"`
	MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
}

func main(){
    faker.SetRandomMapAndSliceSize(20) //Random generated map or array size wont exceed 20...
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}

```
Result:
```
{
    Inta:7
    Int8:-102
    Int16:556
    Int32:113
    Int64:70
    UInta:78
    UInt8:54
    UInt16:1797
    UInt32:8
    UInt64:34
    ASString:[
        geHYIpEoQhQdijFooVEAOyvtTwJOofbQPJdbHvEEdjueZaKIgI
        WVJBBtmrrVccyIydAiLSkMwWbFzFMEotEXsyUXqcmBTVORlkJK
        xYiRTRSZRuGDcMWYoPALVMZgIXoTQtmdGXQfbISKJiavLspuBV
        qsoiYlyRbXLDAMoIdQhgMriODYWCTEYepmjaldWLLjkulDGuQN
        GQXUlqNkVjPKodMebPIeoZZlxfhbQJOjHSRjUTrcgBFPeDZIxn
        MEeRkkLceDsqKLEJFEjJxHtYrYxQMxYcuRHEDGYSPbELDQLSsj
        tWUIACjQWeiUhbboGuuEQIhUJCRSBzVImpYwOlFbsjCRmxboZW
        ZDaAUZEgFKbMJoKIMpymTreeZGWLXNCfVzaEyWNdbkaZOcsfst
        uwlsZBMlEknIBsALpXRaplZWVtXTKzsWglRVBpmfsQfqraiEYA
        AXszbzsOzYPYeXHXRwoPmoPoBxopdFFvWMBTPCxESTepRpjlnB
        kTuOPHlUrSzUQRmZMYplWbyoBbWzQYCiydyzurOduhjuyiGrCE
        FZbeLMbelIeCMnixknIARZRbwALObGXADQqianJbkiEAqqpdnK
        TiQrZbnkvxEciyKXlliUDOGVdpMoAsHSalFbLcYyXxNFLAhqjy
        KlbjbloxkWKSqvUfJQPpFLoddWgeABfYUoaAnylKmEHwxgNsnO
        ]
    SString:VVcaPSFrOPYlEkpVyTRbSZneB
    MSString:map[
        ueFBFTTmqDwrXDoXAYTRhQRmLXhudA:AhQewvZfrlytbAROzGjpXUmNQzIoGl
        fZwrCsFfZwqMsDJXOUYIacflFIeyFU:VMufFCRRHTtuFthOrRAMbzbKVJHnvJ
        rHDQTyZqZVSPLwTtZfNSwKWrgmRghL:lRSXNHkhUyjDuBgoAfrQwOcHYilqRB
        BvCpQJMHzKXKbOoAnTXkLCNxKshwWr:tiNFrXAXUtdywkyygWBrEVrmAcAepD
        uWWKgHKTkUgAZiopAIUmgVWrkrceVy:GuuDNTUiaBtOKwWrMoZDiyaOPxywnq
        HohMjOdMDkAqimKPTgdjUorydpKkly:whAjmraukcZczskqycoJELlMJTghca
        umEgMBGUvBptdKImKsoWXMGJJoRbgT:tPpgHgLEyHmDOocOiSgTbXQHVduLxP
        SRQLHjBXCXKvbLIktdKeLwMnIFOmbi:IJBpLyTcraOxOUtwSKTisjElpulkTL
        dbnDeJZLqMXQGjbTSNxPSlfDHGCghU:JWrymovFwNWbIQBxPpQmlgJsgpXcui
        roraKNGnBXnrJlsxTnFgxHyZeTXdAC:XIcLWqUAQAbfkRrgfjrTVxZCvRJXyl
        TrvxqVVjXAboYDPvUglSJQrltPjzLx:nBhWdfNPybnNnCyQlSshWKOnwUMQzL
        dTHhWJWMwfVvKpIKTFCaoBJgKmnfbD:ixjNHsvSkRkFiNLpgUzIKPsheqhCeY
        lWyBrtfcGWiNbSTJZJXwOPvVngZZMk:kvlYeGgwguVtiafGKjHWsYWewbaXte
        bigsYNfVcNMGtnzgaqEjeRRlIcUdbR:hYOnJupEOvblTTEYzZYPuTVmvTmiit
        ]
    MIint:map[7:7 5:7 8:8 9:5 6:5]
}
```
