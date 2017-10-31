package faker

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bxcodec/faker/support/slice"
)

var phone Phoner

// Constructor
func GetPhoner() Phoner {
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

// 201-886-0269
func (p Phone) PhoneNumber() string {
	randInt, _ := RandomInt(1, 10)
	str := strings.Join(slice.SliceIntToString(randInt), "")
	return fmt.Sprintf("%s-%s-%s", str[:3], str[3:6], str[6:10])
}

// example : (888) 937-7238
func (p Phone) TollFreePhoneNumber() string {
	out := ""
	boxDigitsStart := []string{"777", "888"}

	ints, _ := RandomInt(1, 9)
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
	ints, _ := RandomInt(1, 10)

	for _, v := range slice.SliceIntToString(ints) {
		out += string(v)
	}
	return fmt.Sprintf("+%s%s", boxDigitsStart[rand.Intn(1)], strings.Join(slice.SliceIntToString(ints), ""))
}
