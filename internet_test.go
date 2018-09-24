package faker

import (
	"strings"
	"testing"

	"github.com/bxcodec/faker/support/slice"
)

func TestEmail(t *testing.T) {
	if !strings.Contains(GetNetworker().Email(), "@") {
		t.Error("Expected  email")
	}
}
func TestMacAddress(t *testing.T) {
	i := Internet{}

	if strings.Count(i.MacAddress(), ":") != 5 {
		t.Error("Expected mac address")
	}
}
func TestDomainName(t *testing.T) {
	preTld := strings.Split(GetNetworker().DomainName(), ".")

	if !slice.Contains(tld, preTld[1]) {
		t.Error("Expected get DomainName")
	}
}
func TestURLOneVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/"}

	if !strings.Contains(GetNetworker().URL(), "http") {
		t.Error("Expected get url")
	}
}
func TestURLTwoVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/%s"}

	if !strings.Contains(GetNetworker().URL(), "http") {
		t.Error("Expected get url")
	}
}
func TestUserName(t *testing.T) {
	if GetNetworker().UserName() == "" {
		t.Error("Expected get username")
	}

}
func TestIPv4(t *testing.T) {
	if strings.Count(GetNetworker().IPv4(), ".") != 3 {
		t.Error("Expected IPv4 format")
	}
}
func TestIPv6(t *testing.T) {
	if strings.Count(GetNetworker().IPv6(), ":") != 7 {
		t.Error("Expected IPv4 format")
	}
}
func TestSetNetwork(t *testing.T) {
	SetNetwork(Internet{})
}

func TestPassword(t *testing.T) {
	if GetNetworker().Password() == "" {
		t.Error("Expected hash password")
	}
}
