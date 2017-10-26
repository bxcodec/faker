package faker

import (
	"github.com/agoalofalife/faker/support/slice"
	"strings"
	"fmt"
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
}

type Phone struct{}

func (p Phone) PhoneNumber() string {
	str, _ := RandomInt(100, 999, 3)
	return strings.Join(slice.SliceIntToString(str), "-")
}

// example : (888) 937-7238
func (p Phone) TollFreePhoneNumber() string {
	out := ""
	ints, _ := RandomInt(1, 9, 7)
		for index, v := range slice.SliceIntToString(ints) {
			if	index == 3{
				out += "-"
			}
			out += string(v)
		}

	return fmt.Sprintf("(888) %s", out)
}
