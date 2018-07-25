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
## Example

### With Tag
Supported tag:

**Internet :**
* Email
* Mac address
* Domain name
* Url
* UserName
* IP Address (IPV4 IPV6 )
* Password

 
**Payment :**
* Credit Card Type (VISA, MASTERCARD , AMERICAN EXPRESS ,DISCOVER)
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
* Sentences

```go

package main

import (
	"fmt"
	"github.com/bxcodec/faker"
)

type SomeStruct struct {
 Latitude         float32 `faker:"lat"`
 Long             float32 `faker:"long"`
 CreditCardType   string  `faker:"cc_type"`
 CreditCardNumber string  `faker:"cc_number"`
 Email            string  `faker:"email"`
 IPV4             string  `faker:"ipv4"`
 IPV6             string  `faker:"ipv6"`
 Password         string  `faker:"password"` 
 PhoneNumber      string  `faker:"phone_number"`
 MacAddress       string  `faker:"mac_address"`
 Url              string  `faker:"url"`
 UserName         string  `faker:"username"`
 ToolFreeNumber   string  `faker:"tool_free_number"`
 E164PhoneNumber  string  `faker:"e_164_phone_number"`
 TitleMale        string  `faker:"title_male"`
 TitleFemale      string  `faker:"title_female"`
 FirstNameMale    string  `faker:"first_name_male"`
 FirstNameFemale  string  `faker:"first_name_female"`
 LastName         string  `faker:"last_name"`
 Name             string  `faker:"name"`
 UnixTime         int64   `faker:"unix_time"`
 Date             string  `faker:"date"`
 Time             string  `faker:"time"`
 MonthName        string  `faker:"month_name"`
 Year             string  `faker:"year"`
 DayOfWeek        string  `faker:"day_of_week"`
 DayOfMonth       string  `faker:"day_of_month"`
 Timestamp        string  `faker:"timestamp"`
 Century          string  `faker:"century"`
 TimeZone         string  `faker:"timezone"`
 TimePeriod       string  `faker:"time_period"`
 Word             string  `faker:"word"`
 Sentence         string  `faker:"sentence"`
 Sentences        string  `faker:"sentences"`
}

func main() {

  a= SomeStruct{}
  err:= faker.FakeData(&a)
  if err!= nil {
    fmt.Println(err)
  }
  fmt.Printf("%+v", a)
	//Will Print  :
	 /* Latitude: -74.209991,
        Long: -4.394531,
        CreditCardNumber: 376243757700558,
        CreditCardType: american express,
        Email: QjTIcmn@lmErh.info,
        IPV4: 80.206.109.93,
        IPV6: 448f:a391:90c6:d165:1d1f:b536:45f7:f084,
	    Password: FhOSWpnzXwxxnCTlZJcZeBTUeRaVidThUdntQryyjNiAFwbyCe,
        PhoneNumber: 865-321-1047,
        MacAddress: 46:77:6f:bb:d0:fc,
        Url: https://sFUBEIj.info/kYvcpYL.html,
        UserName: VwrVhzZ,
        ToolFreeNumber: (777) 372-148965,
        E164PhoneNumber: +765101283947,
        TitleMale: Prof.,
        TitleFemale: Dr.,
        FirstNameMale: Charley,
        FirstNameFemale: Freeda,
        LastName: Runte,
        Name: Dr. Freeda Runte,
        UnixTime: 679631083,
        Date: 2010-06-12,
        Time: 11:33:24,
        MonthName: January,
        Year: 1982,
        DayOfWeek: Sunday,
        DayOfMonth: 26,
        Timestamp: 1972-06-24 14:03:32,
        Century: XV,
        TimeZone: Navajo,
        TimePeriod: AM,
	    Word: consectetur,
        Sentence: Perferendis aut sit accusantium consequatur voluptatem .,
        Sentences: Perferendis sit consequatur accusantium aut voluptatem .Aut voluptatem consequatur sit perferendis accusantium .Aut consequatur sit accusantium perferendis voluptatem .,
	 */
}

```

### Without Tag

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
  err:= faker.FakeData(&a)
  if err!= nil {
    fmt.Println(err)
  }
  fmt.Printf("%+v", a)
}
```
Output :

```bash
{Int:7088293148785081331 Int8:7 Int16:14 Int32:1777976883 Int64:2467854463682814928 String:XMhCTmwvVqEUryIKnpWrQmBdb Bool:true SString:[iiCGZ GESVVaP] SInt:[2391903971675293806 5270400206229440165 7315288441301820955] SInt8:[124 104 84] SInt16:[-9403 -23327 -19174] SInt32:[1714966339 1617248797 1233525792] SInt64:[6505581000035730776 989945489581839946 7467254172609770414] SFloat32:[0.6761954 0.13427323 0.35608092] SFloat64:[0.49714054026277343 0.29188223737765046 0.7800285978504301] SBool:[true true true] Struct:{Number:8662858647992239649 Height:2466984558238338402 AnotherStruct:{Image:kNIwoxPiVcOqQxBUyyAuDAKom}}}
```

![Example to use Faker](https://cdn-images-1.medium.com/max/800/1*AkMbxngg7zfvtWiuvFb4Mg.gif)

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
* Not support for private field. Just make sure your field's struct is public. If not, it will throw panic error.
* Not support for `interface{}` data type. How we can generate if we don't know what is the data type? 
* Not support for `map[interface{}]interface{}, map[any_type]interface{}, map[interface{}]any_type`. Still, it's about interface. We can't give you something if we don't know what really you want. 
* Not fully support for custom type, but a few custom type already supported, still investigating how to do this in the correct ways. For now, if you use `faker`, it's safer not to use any custom type to avoid panic. 

## Contribution
To contrib on this project, you can make a PR or just an issue.

### Maintainer
- <a href="https://github.com/bxcodec">  **Iman Tumorang** </a> <br> 
- <a href="https://github.com/agoalofalife">  **Ilya** </a> <br> 


