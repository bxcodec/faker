package faker

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestDataFaker(t *testing.T) {
	SetDataFaker(Lorem{})
}

func TestWord(t *testing.T) {
	word, err := GetLorem().Word(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !slice.Contains(wordList, word.(string)) {
		t.Error("Expected word from slice wordList")
	}
}

func TestSentence(t *testing.T) {
	res, err := GetLorem().Sentence(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	s := res.(string)
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected sentence")
	}
}

func TestParagraph(t *testing.T) {
	res, err := GetLorem().Paragraph(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	s := res.(string)
	fmt.Println(s)
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected paragraph")
	}
}
