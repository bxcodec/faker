package faker

import (
	"fmt"
	"github.com/agoalofalife/faker/support/slice"
	"math/rand"
	"strings"

)

var phone Phoner

// Constructor
func getPhoner() Phoner {
	mu.Lock()
	defer mu.Unlock()

	if phone == nil {
		phone = &Phone{}
	}
	return phone
}

// this set custom Phoner
func SetPhoner(p Phoner) {
	phone = p
}

type Phoner interface {
	PhoneNumber() string
	TollFreePhoneNumber() string
	E164PhoneNumber() string
}

type Phone struct{}

func (p Phone) PhoneNumber() string {
	str, _ := RandomInt(100, 999, 3)
	return strings.Join(slice.SliceIntToString(str), "-")
}

// example : (888) 937-7238
func (p Phone) TollFreePhoneNumber() string {
	out := ""
	boxDigitsStart := []string{"777", "888"}

	ints, _ := RandomInt(1, 9, 10)
	for index, v := range slice.SliceIntToString(ints) {
		if index == 3 {
			out += "-"
		}
		out += string(v)
	}
	return fmt.Sprintf("(%s) %s", boxDigitsStart[rand.Intn(1)], out)
}

// '+27113456789'
func (p Phone) E164PhoneNumber() string {
	out := ""
	boxDigitsStart := []string{"7", "8"}
	ints, _ := RandomInt(1, 9, 10)

	for _, v := range slice.SliceIntToString(ints) {
		out += string(v)
	}
	return fmt.Sprintf("+%s%s", boxDigitsStart[rand.Intn(1)], strings.Join(slice.SliceIntToString(ints), ""))
}
