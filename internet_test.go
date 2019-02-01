package faker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3/support/slice"
)

func TestEmail(t *testing.T) {
	email, err := GetNetworker().Email(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.Contains(email.(string), "@") {
		t.Error("Expected  email")
	}
}
func TestMacAddress(t *testing.T) {
	i := Internet{}
	mc, err := i.MacAddress(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}

	if strings.Count(mc.(string), ":") != 5 {
		t.Error("Expected mac address")
	}
}
func TestDomainName(t *testing.T) {
	domain, err := GetNetworker().DomainName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}

	preTld := strings.Split(domain.(string), ".")

	if !slice.Contains(tld, preTld[1]) {
		t.Error("Expected get DomainName")
	}
}
func TestURLOneVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/"}
	res, err := GetNetworker().URL(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.Contains(res.(string), "http") {
		t.Error("Expected get url")
	}
}
func TestURLTwoVerbs(t *testing.T) {
	urlFormats = []string{
		"http://www.%s/%s"}
	res, err := GetNetworker().URL(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.Contains(res.(string), "http") {
		t.Error("Expected get url")
	}
}
func TestUserName(t *testing.T) {
	usrname, err := GetNetworker().UserName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if usrname.(string) == "" {
		t.Error("Expected get username")
	}

}
func TestIPv4(t *testing.T) {
	ip, err := GetNetworker().IPv4(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if strings.Count(ip.(string), ".") != 3 {
		t.Error("Expected IPv4 format")
	}
}
func TestIPv6(t *testing.T) {
	ip, err := GetNetworker().IPv6(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if strings.Count(ip.(string), ":") != 7 {
		t.Error("Expected IPv4 format")
	}
}
func TestSetNetwork(t *testing.T) {

	SetNetwork(Internet{})
}

func TestPassword(t *testing.T) {
	pass, err := GetNetworker().Password(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if pass.(string) == "" {
		t.Error("Expected hash password")
	}
}
