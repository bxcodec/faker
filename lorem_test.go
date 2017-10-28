package faker

import (
	"github.com/agoalofalife/faker/support/slice"
	"testing"
)

func TestDataFaker(t *testing.T) {
	SetDataFaker(Lorem{})
}

func TestWord(t *testing.T) {
	if !slice.Contains(wordList, getLorem().Word()) {
		t.Error("Expected word from slice wordList")
	}
}
