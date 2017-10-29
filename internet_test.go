package faker

import (
	"github.com/bxcodec/faker/support/slice"
	"strings"
	"testing"
)

func TestEmail(t *testing.T) {
	if !strings.Contains(getNetworker().Email(), "@") {
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
	preTld := strings.Split(getNetworker().DomainName(), ".")

	if !slice.Contains(tld, preTld[1]) {
		t.Error("Expected get DomainName")
	}
}
func TestUrlOneVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/"}

	if !strings.Contains(getNetworker().Url(), "http") {
		t.Error("Expected get url")
	}
}
func TestUrlTwoVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/%s"}

	if !strings.Contains(getNetworker().Url(), "http") {
		t.Error("Expected get url")
	}
}
func TestUserName(t *testing.T) {
	if getNetworker().UserName() == "" {
		t.Error("Expected get username")
	}

}
func TestIpv4(t *testing.T) {
	if strings.Count(getNetworker().Ipv4(), ".") != 3 {
		t.Error("Expected Ipv4 format")
	}
}
func TestIpv6(t *testing.T) {
	if strings.Count(getNetworker().Ipv6(), ":") != 7 {
		t.Error("Expected Ipv4 format")
	}
}
func TestSetNetwork(t *testing.T) {
	SetNetwork(Internet{})
}

func TestPassword(t *testing.T) {
	if getNetworker().Password() == "" {
		t.Error("Expected hash password")
	}
}
