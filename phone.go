package faker

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/bxcodec/faker/v3/support/slice"
)

var phone Phoner

var usaAreaCodes = [52]string{
	"518", "410", "404", "207", "512",
	"225", "701", "208", "617", "775",
	"843", "307", "803", "614", "603",
	"303", "515", "302", "502", "717",
	"860", "406", "808", "317", "601",
	"904", "573", "907", "517", "402",
	"501", "608", "334", "802", "615",
	"405", "360", "602", "605", "401",
	"919", "804", "916", "651", "503",
	"385", "505", "417", "850", "785",
	"609", "202"}

// GetPhoner serves as a constructor for Phoner interface
func GetPhoner() Phoner {
	mu.Lock()
	defer mu.Unlock()

	if phone == nil {
		phone = &Phone{}
	}
	return phone
}

// SetPhoner sets custom Phoner
func SetPhoner(p Phoner) {
	phone = p
}

// Phoner serves overall tele-phonic contact generator
type Phoner interface {
	PhoneNumber(v reflect.Value) (interface{}, error)
	TollFreePhoneNumber(v reflect.Value) (interface{}, error)
	E164PhoneNumber(v reflect.Value) (interface{}, error)
}

// Phone struct
type Phone struct {
}

func (p Phone) phonenumber() string {
	randInt, _ := RandomInt(1, 4)
	str := strings.Join(slice.IntToString(randInt), "")
	return fmt.Sprintf("%s-555-%s", usaAreaCodes[rand.Intn(len(usaAreaCodes))], str)
}

// PhoneNumber generates phone numbers of type: "201-886-0269"
func (p Phone) PhoneNumber(v reflect.Value) (interface{}, error) {
	return p.phonenumber(), nil
}

func (p Phone) tollfreephonenumber() string {
	out := ""
	boxDigitsStart := []string{"777", "888"}

	ints, _ := RandomInt(1, 9)
	for index, v := range slice.IntToString(ints) {
		if index == 3 {
			out += "-"
		}
		out += v
	}
	return fmt.Sprintf("(%s) %s", boxDigitsStart[rand.Intn(1)], out)
}

// TollFreePhoneNumber generates phone numbers of type: "(888) 937-7238"
func (p Phone) TollFreePhoneNumber(v reflect.Value) (interface{}, error) {
	return p.tollfreephonenumber(), nil
}

func (p Phone) e164PhoneNumber() string {
	out := ""
	boxDigitsStart := []string{"7", "8"}
	ints, _ := RandomInt(1, 10)

	for _, v := range slice.IntToString(ints) {
		out += v
	}
	return fmt.Sprintf("+%s%s", boxDigitsStart[rand.Intn(1)], strings.Join(slice.IntToString(ints), ""))
}

// E164PhoneNumber generates phone numbers of type: "+27113456789"
func (p Phone) E164PhoneNumber(v reflect.Value) (interface{}, error) {
	return p.e164PhoneNumber(), nil
}
