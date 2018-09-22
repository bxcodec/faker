package faker

import (
	"strings"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestDataFaker(t *testing.T) {
	SetDataFaker(Lorem{})
}

func TestWord(t *testing.T) {
	if !slice.Contains(wordList, GetLorem().Word()) {
		t.Error("Expected word from slice wordList")
	}
}

func TestSentence(t *testing.T) {
	s := GetLorem().Sentence()
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected sentence")
	}
}

func TestParagraph(t *testing.T) {
	s := GetLorem().Paragraph()
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected paragraph")
	}
}
