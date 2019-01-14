# Docs

## [faker](#)

Struct Data Fake Generator

Faker  will generate you a fake data based on your Struct.

[![Build Status](https://travis-ci.org/bxcodec/faker.svg?branch=master)](https://travis-ci.org/bxcodec/faker)
[![codecov](https://codecov.io/gh/bxcodec/faker/branch/master/graph/badge.svg)](https://codecov.io/gh/bxcodec/faker)
[![Go Report Card](https://goreportcard.com/badge/github.com/bxcodec/faker)](https://goreportcard.com/report/github.com/bxcodec/faker)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/faker/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/bxcodec/faker?status.svg)](https://godoc.org/github.com/bxcodec/faker)

## Index

* [Support](#support)
* [Getting Started](#getting-started)
* [Example](#example)
* [Limitation](#limitation)
* [Contribution](#contribution)


## Support

You can file an [Issue](https://github.com/bxcodec/faker/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/bxcodec/faker)


## Getting Started

#### Download

```shell
go get -u github.com/bxcodec/faker
```
# Example

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
* Credit Card Type (VISA, MASTERCARD, AMERICAN EXPRESS, DISCOVER)
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
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	ToolFreeNumber     string  `faker:"tool_free_number"`
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
			IPV4: 99.23.42.63
			IPV6: 975c:fb2c:2133:fbdd:beda:282e:1e0a:ec7d
			Password: dfJdyHGuVkHBgnHLQQgpINApynzexnRpgIKBpiIjpTPOmNyMFb
			PhoneNumber: 792-153-4861
			MacAddress: cd:65:e1:d4:76:c6
			URL: https://www.oEuqqAY.org/QgqfOhd
			UserName: lVxELHS
			ToolFreeNumber: (777) 831-964572
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

## Custom Generator Provider
You can also add your own generator function to your own defined tags. See example below
```go
type Gondoruwo struct {
	Name       string
	Locatadata int
}

type Sample struct {
	ID                 int64     `faker:"customIdFaker"`
	Gondoruwo          Gondoruwo `faker:"gondoruwo"`
	Danger             string    `faker:"danger"`
}

func CustomGenerator() {
	faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		 return int64(43), nil
	})
	faker.AddProvider("danger", func(v reflect.Value) (interface{}, error) {
		return "danger-ranger", nil
	})

	faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
		obj := Gondoruwo{
			Name:       "Power",
			Locatadata: 324,
		}
		return obj, nil
	})
}

func main() { 
	CustomGenerator()
	var sample Sample
	faker.FakeData(&sample)
	fmt.Printf("%+v", sample)
}
```

Results:
```
{ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger}
```

## Without Tag
You also can use faker to generate your structs data randomly without any tag. And it will fill the data based on its data-type.

```go

package main

import (
	"fmt"

	"github.com/bxcodec/faker"
)

type SomeStruct struct {
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	String   string
	Bool     bool
	SString  []string
	SInt     []int
	SInt8    []int8
	SInt16   []int16
	SInt32   []int32
	SInt64   []int64
	SFloat32 []float32
	SFloat64 []float64
	SBool    []bool
	Struct   AStruct
}
type AStruct struct {
	Number        int64
	Height        int64
	AnotherStruct BStruct
}

type BStruct struct {
	Image string
}

func main() {

	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
	/*
		Output:
		{
		    Int:5231564546548329
		    Int8:52
		    Int16:8
		    Int32:2046951236
		    Int64:1486554863682234423
		    String:ELIQhSfhHmWxyzRPOTNblEQsp
		    Bool:false
		    SString:[bzYwplRGUAXPwatnpVMWSYjep zmeuJVGHHgmIsuyWmLJnDmbTI FqtejCwoDyMBWatoykIzorCJZ]
		    SInt:[11661230973193 626062851427 12674621422454 5566279673347]
		    SInt8:[12 2 58 22 11 66 5 88]
		    SInt16:[29295225 8411281 69902706328]
		    SInt32:[60525685140 2733640366211 278043484637 5167734561481]
		    SInt64:[81684520429188374184 9917955420365482658170 996818723778286568 163646873275501565]
		    SFloat32:[0.556428 0.7692596 0.6779895 0.29171365 0.95445055]
		    SFloat64:[0.44829454895586585 0.5495675898536803 0.6584538253883265]
		    SBool:[true false true false true true false]
		    Struct:{
		        Number:1
		        Height:26
		        AnotherStruct:{
		            Image:RmmaaHkAkrWHldVbBNuQSKlRb
		        }
		    }
		}
	*/
}

```

## DEMO

![Example to use Faker](https://cdn-images-1.medium.com/max/800/1*AkMbxngg7zfvtWiuvFb4Mg.gif)

## Benchmark
Bench To Generate Fake Data
#### Without Tag
```bash
BenchmarkFakerDataNOTTagged-4             500000              3049 ns/op             488 B/op         20 allocs/op
```

#### Using Tag
```bash
 BenchmarkFakerDataTagged-4                100000             17470 ns/op             380 B/op         26 allocs/op
```

### MUST KNOW
The Struct Field must PUBLIC.<br>
Support Only For :
* int  int8  int16  int32  int64
* []int  []int8  []int16  []int32  []int64
* bool []bool
* string []string
* float32 float64 []float32 []float64
* Nested Struct Field
* time.Time []time.Time

## Limitation
Unfortunately this library has some limitation
* Not support for private field. Just make sure your field's struct is public. If not, it will throw panic error. You can avoid panic using tag skip `faker:"-"` in the private field.
* Not support for `interface{}` data type. How we can generate if we don't know what is the data type?
* Not support for `map[interface{}]interface{}, map[any_type]interface{}, map[interface{}]any_type`. Still, it's about interface. We can't give you something if we don't know what really you want.
* Not fully support for custom type, but a few custom type already supported, still investigating how to do this in the correct ways. For now, if you use `faker`, it's safer not to use any custom type to avoid panic.

## Contribution
To contrib on this project, you can make a PR or just an issue.

### Maintainer
- <a href="https://github.com/bxcodec">  **Iman Tumorang** </a> <br>
- <a href="https://github.com/agoalofalife">  **Ilya** </a> <br>


