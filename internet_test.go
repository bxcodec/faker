package faker

import (
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/pkg/slice"
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

func TestPassword(t *testing.T) {
	pass, err := GetNetworker().Password(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if pass.(string) == "" {
		t.Error("Expected hash password")
	}
}

func TestFakeEmail(t *testing.T) {
	email := Email()
	if !strings.Contains(email, "@") {
		t.Error("Expected  email")
	}
}
func TestFakeMacAddress(t *testing.T) {
	mc := MacAddress()
	if strings.Count(mc, ":") != 5 {
		t.Error("Expected mac address")
	}
}
func TestFakeDomainName(t *testing.T) {
	domain := DomainName()
	preTld := strings.Split(domain, ".")

	if !slice.Contains(tld, preTld[1]) {
		t.Error("Expected get DomainName")
	}
}
func TestFakeURL(t *testing.T) {
	resURL := URL()
	if !strings.Contains(resURL, "http") {
		t.Error("Expected get url")
	}
}

func TestFakeUserName(t *testing.T) {
	usrname := Username()
	if usrname == "" {
		t.Error("Expected get username")
	}
}
func TestFakeIPv4(t *testing.T) {
	ip := IPv4()
	if strings.Count(ip, ".") != 3 {
		t.Error("Expected IPv4 format")
	}
}
func TestFakeIPv6(t *testing.T) {
	ip := IPv6()
	if strings.Count(ip, ":") != 7 {
		t.Error("Expected IPv4 format")
	}
}

func TestFakePassword(t *testing.T) {
	pass := Password()
	if pass == "" {
		t.Error("Expected hash password")
	}
}

func TestFakeJWT(t *testing.T) {
	jwt := Jwt()
	reg := regexp.MustCompile(`[a-zA-Z]+.[a-zA-Z]+.[a-zA-Z]+`)
	if !reg.MatchString(jwt) {
		t.Error("Invalid format on JWT token")
	}
}
