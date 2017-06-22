# Docs

## [faker](#)

Struct Data Fake Generator

Faker  will generate you a fake data based on your Struct.


[![License](https://img.shields.io/badge/status-on%20going-yellowgreen.svg)](#)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/bxcodec/saint/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/bxcodec/saint?status.svg)](https://godoc.org/github.com/bxcodec/saint)

## Index

* [Support](#support)
* [Getting Started](#getting-started)
* [Example](#example)


## Support


You can also email <iman.tumorang@gmail.com> or file an [Issue](https://github.com/bxcodec/faker/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/bxcodec/faker)




## Getting Started

#### Download

```shell
go get -u github.com/bxcodec/faker
```
## Example

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
	Number int64
	Height int64
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
{Int:8906957488773767119 Int8:6 Int16:14 Int32:391219825 Int64:2374447092794071106 String:poraKzAxVbWVkMkpcZCcWlYMd Bool:false SString:[MehdV aVotHsi] SInt:[528955241289647236 7620047312653801973 2774096449863851732] SInt8:[122 -92 -92] SInt16:[15679 -19444 -30246] SInt32:[1146660378 946021799 852909987] SInt64:[6079203475736033758 6913211867841842836 3269201978513619428] SFloat32:[0.019562425 0.12729558 0.36450312] SFloat64:[0.7825838989890364 0.9732903338838912 0.8316541489234004] SBool:[true false true] Struct:{Number:7693944638490551161 Height:6513508020379591917}}
```


Bench To Generate Fake Data
```bash
BenchmarkFakerData-4      300000              4653 ns/op             880 B/op         31 allocs/op
```
