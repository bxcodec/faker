package faker

import (
	"testing"
	"strings"
)

func TestUrl(t *testing.T)  {
	i := Internet{}

	if strings.Contains(i.Url(), "http") == false {
		t.Error("Expected get url")
	}
}

func TestMacAddress(t *testing.T)  {
	i := Internet{}

	if 	strings.Count(i.MacAddress(), ":") != 5 {
		t.Error("Expected mac address")
	}
}