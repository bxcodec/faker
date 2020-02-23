package faker_test

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// SomeStruct ...
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

// AStruct ...
type AStruct struct {
	Number        int64
	Height        int64
	AnotherStruct BStruct
}

// BStruct ...
type BStruct struct {
	Image string
}

// You also can use faker to generate your structs data randomly without any tag.
// And it will fill the data based on its data-type.
func Example_withoutTag() {
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
	/*
		Result:
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
