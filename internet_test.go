package faker

import (
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v4/options"
	"github.com/bxcodec/faker/v4/support/slice"
)

func TestEmail(t *testing.T) {
	email, err := GetNetworker(*options.DefaultOption()).Email(reflect.Value{})
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
	domain, err := GetNetworker(*options.DefaultOption()).DomainName(reflect.Value{})
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
	res, err := GetNetworker(*options.DefaultOption()).URL(reflect.Value{})
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
	res, err := GetNetworker(*options.DefaultOption()).URL(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !strings.Contains(res.(string), "http") {
		t.Error("Expected get url")
	}
}
func TestUserName(t *testing.T) {
	usrname, err := GetNetworker(*options.DefaultOption()).UserName(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if usrname.(string) == "" {
		t.Error("Expected get username")
	}

}
func TestIPv4(t *testing.T) {
	ip, err := GetNetworker(*options.DefaultOption()).IPv4(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if strings.Count(ip.(string), ".") != 3 {
		t.Error("Expected IPv4 format")
	}
}
func TestIPv6(t *testing.T) {
	ip, err := GetNetworker(*options.DefaultOption()).IPv6(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if strings.Count(ip.(string), ":") != 7 {
		t.Error("Expected IPv4 format")
	}
}

func TestPassword(t *testing.T) {
	pass, err := GetNetworker(*options.DefaultOption()).Password(reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if pass.(string) == "" {
		t.Error("Expected hash password")
	}
}

func TestFakeEmail(t *testing.T) {
	email := Email(options.DefaultOption())
	if !strings.Contains(email, "@") {
		t.Error("Expected  email")
	}
}
func TestFakeMacAddress(t *testing.T) {
	mc := MacAddress(options.DefaultOption())
	if strings.Count(mc, ":") != 5 {
		t.Error("Expected mac address")
	}
}
func TestFakeDomainName(t *testing.T) {
	domain := DomainName(options.DefaultOption())
	preTld := strings.Split(domain, ".")

	if !slice.Contains(tld, preTld[1]) {
		t.Error("Expected get DomainName")
	}
}
func TestFakeURL(t *testing.T) {
	resURL := URL(options.DefaultOption())
	if !strings.Contains(resURL, "http") {
		t.Error("Expected get url")
	}
}

func TestFakeUserName(t *testing.T) {
	usrname := Username(options.DefaultOption())
	if usrname == "" {
		t.Error("Expected get username")
	}
}
func TestFakeIPv4(t *testing.T) {
	ip := IPv4(options.DefaultOption())
	if strings.Count(ip, ".") != 3 {
		t.Error("Expected IPv4 format")
	}
}
func TestFakeIPv6(t *testing.T) {
	ip := IPv6(options.DefaultOption())
	if strings.Count(ip, ":") != 7 {
		t.Error("Expected IPv4 format")
	}
}

func TestFakePassword(t *testing.T) {
	pass := Password(options.DefaultOption())
	if pass == "" {
		t.Error("Expected hash password")
	}
}

func TestFakeJWT(t *testing.T) {
	jwt := Jwt(options.DefaultOption())
	reg := regexp.MustCompile(`[a-zA-Z]+.[a-zA-Z]+.[a-zA-Z]+`)
	if !reg.MatchString(jwt) {
		t.Error("Invalid format on JWT token")
	}
}
