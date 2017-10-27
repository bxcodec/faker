package faker

import (
	"testing"
	"github.com/bxcodec/faker/support/slice"
)
func TestTitleMale(t *testing.T) {
	p := getPerson()
	if !slice.Contains(titleMales, p.TitleMale()) {
		t.Error("Expected value from variable titleMales in function TitleMale")
	}

}