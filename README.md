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
* [Contribution](#contributions)


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

}

func main() {

  a= SomeStruct{}
  err:= faker.FakeData(&a)
  if err!= nil {
    fmt.Println(err)
  }
  fmt.Printf("%+v", a)
	//Will Print  :
	// {Latitude:31.456718 Long:159.9867 CreditCardType:Discover CreditCardNumber:6011091892846730 Email:MeHOZEM@Zxhmd.com IPV4:76.24.239.144 IPV6:e8e9:4a26:4c2b:20d3:541e:98b:c5fd:aa26}PASS

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

  a= SomeStruct{}
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
```bash
BenchmarkFakerData-4      300000              4653 ns/op             880 B/op         31 allocs/op
```

### MUST KNOW
The Struct Field must PUBLIC.<br>
Support Only For :
* int  int8  int16  int32  int64
* []int  []int8  []int16  []int32  []int64  
* bool []bool
* string []string
* float32 float64 []float32 []float64
* Nested Struct Field for Non POINTER
* time.Time []time.Time
