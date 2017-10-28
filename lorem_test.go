package faker

import (
	"github.com/agoalofalife/faker/support/slice"
	"strings"
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

func TestSentence(t *testing.T) {
	s := getLorem().Sentence()
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected sentence")
	}
}
