package faker

type Dowser interface {
	TitleMale() string
}

var person Dowser
var titleMales = []string{
	"Mr.", "Dr.", "Prof.",
}
func getPerson() Dowser {
	mu.Lock()
	defer mu.Unlock()

	if person == nil {
		person = &Person{}
	}
	return person
}

// this set custom Address
func SetDowser(d Dowser) {
	person = d
}

type Person struct {
	
}

func (p Person) TitleMale() string {
	return randomElementFromSliceString(titleMales)
}