package slice

import (
	"strconv"
)

// Check item in slice string type
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// Convert slice int to slice string
func SliceIntToString(intSl []int) (str []string) {
	for i := range intSl {
		number := intSl[i]
		text := strconv.Itoa(number)
		str = append(str, text)
	}
	return str
}
