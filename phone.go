package faker

import (
	"github.com/agoalofalife/faker/support/slice"
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
}

type Phone struct{}

func (p Phone) PhoneNumber() string {
	str, _ := RandomInt(100, 999, 3)
	return strings.Join(slice.SliceIntToString(str), "-")
}
