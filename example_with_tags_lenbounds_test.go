package faker_test

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// You can set length for your random strings also set boundary for your integers.
func Example_withTagsLengthAndBoundary() {
	// SomeStruct ...
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

		Float32 float32 `faker:"boundary_start=12.65, boundary_end=184.05"`
		Float64 float64 `faker:"boundary_start=1.256, boundary_end=3.4"`

		ASString []string          `faker:"len=50"`
		SString  string            `faker:"len=25"`
		MSString map[string]string `faker:"len=30"`
		MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
	}

	_ = faker.SetRandomMapAndSliceSize(20) // Random generated map or array size wont exceed 20...
	a := SomeStruct{}
	_ = faker.FakeData(&a)
	fmt.Printf("%+v", a)
	// Result:
	/*
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
	       Float32:60.999058
	       Float64:2.590148738554016
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
	*/
}
