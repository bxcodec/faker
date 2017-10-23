package faker

import (
	"testing"
	"strings"
	"log"
)

func TestUrl(t *testing.T)  {
	i := Internet{}

	if strings.Contains(i.Url(), "http") == false {
		t.Error("Expected get url")
	}
}

func TestMacAddress(t *testing.T)  {
	i := Internet{}

	log.Println(i.MacAddress())
}