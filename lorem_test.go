package faker

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/pkg/options"
	"github.com/bxcodec/faker/v4/pkg/slice"
)

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

func TestFakeWord(t *testing.T) {
	word := Word()
	if !slice.Contains(wordList, word) {
		t.Error("Expected word from slice wordList")
	}
}

func TestFakeSentence(t *testing.T) {
	res := Sentence()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected sentence")
	}
}

func TestFakeParagraph(t *testing.T) {
	res := Paragraph()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected paragraph")
	}
}

func TestUniqueWord(t *testing.T) {
	word := Word(options.WithGenerateUniqueValues(true))
	ResetUnique()
	if !slice.Contains(wordList, word) {
		t.Error("Expected word from slice wordList")
	}
}

func TestUniqueWordPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but didn't")
		}
		ResetUnique()
	}()

	length := len(wordList) + 1
	for i := 0; i < length; i++ {
		Word(options.WithGenerateUniqueValues(true))
	}
	ResetUnique()
}

func TestUniqueSentence(t *testing.T) {
	res := Sentence(options.WithGenerateUniqueValues(true))
	ResetUnique()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected sentence")
	}
}

func TestUniqueParagraph(t *testing.T) {
	res := Paragraph(options.WithGenerateUniqueValues(true))
	ResetUnique()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected paragraph")
	}
}
