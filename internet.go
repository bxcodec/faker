package faker

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/bxcodec/faker/v4/pkg/options"
)

var tld = []string{"com", "biz", "info", "net", "org", "ru"}
var urlFormats = []string{
	"http://www.%s/",
	"https://www.%s/",
	"http://%s/",
	"https://%s/",
	"http://www.%s/%s",
	"https://www.%s/%s",
	"http://%s/%s",
	"https://%s/%s",
	"http://%s/%s.html",
	"https://%s/%s.html",
	"http://%s/%s.php",
	"https://%s/%s.php",
}

// GetNetworker returns a new Networker interface of Internet
func GetNetworker(opts ...options.OptionFunc) Networker {
	opt := options.BuildOptions(opts)
	internet := &Internet{
		fakerOption: *opt,
	}
	return internet
}

// Networker is logical layer for Internet
type Networker interface {
	Email(v reflect.Value) (interface{}, error)
	MacAddress(v reflect.Value) (interface{}, error)
	DomainName(v reflect.Value) (interface{}, error)
	URL(v reflect.Value) (interface{}, error)
	UserName(v reflect.Value) (interface{}, error)
	IPv4(v reflect.Value) (interface{}, error)
	IPv6(v reflect.Value) (interface{}, error)
	Password(v reflect.Value) (interface{}, error)
	Jwt(v reflect.Value) (interface{}, error)
}

// Internet struct
type Internet struct {
	fakerOption options.Options
}

func (internet Internet) email() (string, error) {
	var err error
	var emailName, emailDomain string
	if emailName, err = randomString(7, internet.fakerOption); err != nil {
		return "", err
	}
	if emailDomain, err = randomString(7, internet.fakerOption); err != nil {
		return "", err
	}
	return (emailName + "@" + emailDomain + "." + randomElementFromSliceString(tld)), nil
}

// Email generates random email id
func (internet Internet) Email(v reflect.Value) (interface{}, error) {
	return internet.email()
}

// Email get email randomly in string
func Email(opts ...options.OptionFunc) string {
	return singleFakeData(EmailTag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		r, err := i.email()
		if err != nil {
			panic(err.Error())
		}
		return r
	}, opts...).(string)
}

func (internet Internet) macAddress() string {
	ip := make([]byte, 6)
	for i := 0; i < 6; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.HardwareAddr(ip).String()
}

// MacAddress generates random MacAddress
func (internet Internet) MacAddress(v reflect.Value) (interface{}, error) {
	return internet.macAddress(), nil
}

// MacAddress get mac address randomly in string
func MacAddress(opts ...options.OptionFunc) string {
	return singleFakeData(MacAddressTag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		return i.macAddress()
	}, opts...).(string)
}

func (internet Internet) domainName() (string, error) {
	domainPart, err := randomString(7, internet.fakerOption)
	if err != nil {
		return "", err
	}
	return (domainPart + "." + randomElementFromSliceString(tld)), nil
}

// DomainName generates random domain name
func (internet Internet) DomainName(v reflect.Value) (interface{}, error) {
	return internet.domainName()
}

// DomainName get email domain name in string
func DomainName(opts ...options.OptionFunc) string {
	return singleFakeData(DomainNameTag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		d, err := i.domainName()
		if err != nil {
			panic(err.Error())
		}
		return d
	}, opts...).(string)
}

func (internet Internet) url() (string, error) {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	d, err := internet.domainName()
	if err != nil {
		return "", nil
	}
	if countVerbs == 1 {
		return fmt.Sprintf(format, d), nil
	}
	u, err := internet.username()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(format, d, u), nil
}

// URL generates random URL standardized in urlFormats const
func (internet Internet) URL(v reflect.Value) (interface{}, error) {
	return internet.url()
}

// URL get Url randomly in string
func URL(opts ...options.OptionFunc) string {
	return singleFakeData(URLTag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		u, err := i.url()
		if err != nil {
			panic(err.Error())
		}
		return u
	}, opts...).(string)
}

func (internet Internet) username() (string, error) {
	return randomString(7, internet.fakerOption)
}

// UserName generates random username
func (internet Internet) UserName(v reflect.Value) (interface{}, error) {
	return internet.username()
}

// Username get username randomly in string
func Username(opts ...options.OptionFunc) string {
	return singleFakeData(UserNameTag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		u, err := i.username()
		if err != nil {
			panic(err.Error())
		}
		return u
	}, opts...).(string)
}

func (internet Internet) ipv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To4().String()
}

// IPv4 generates random IPv4 address
func (internet Internet) IPv4(v reflect.Value) (interface{}, error) {
	return internet.ipv4(), nil
}

// IPv4 get IPv4 randomly in string
func IPv4(opts ...options.OptionFunc) string {
	return singleFakeData(IPV4Tag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		return i.ipv4()
	}, opts...).(string)
}

func (internet Internet) ipv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// IPv6 generates random IPv6 address
func (internet Internet) IPv6(v reflect.Value) (interface{}, error) {
	return internet.ipv6(), nil
}

// IPv6 get IPv6 randomly in string
func IPv6(opts ...options.OptionFunc) string {
	return singleFakeData(IPV6Tag, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		return i.ipv6()
	}, opts...).(string)
}

func (internet Internet) password() (string, error) {
	return randomString(50, internet.fakerOption)
}

// Password returns a hashed password
func (internet Internet) Password(v reflect.Value) (interface{}, error) {
	return internet.password()
}

// Password get password randomly in string
func Password(opts ...options.OptionFunc) string {
	return singleFakeData(PASSWORD, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		p, err := i.password()
		if err != nil {
			panic(err.Error())
		}
		return p
	}, opts...).(string)
}

func (internet Internet) jwt() (string, error) {
	element, err := randomString(40, internet.fakerOption)
	sl := element[:]
	if err != nil {
		return "", err
	}
	return strings.Join([]string{sl, sl, sl}, "."), nil
}

// Jwt returns a jwt-like random string in xxxx.yyyy.zzzz style
func (internet Internet) Jwt(v reflect.Value) (interface{}, error) {
	return internet.jwt()
}

// Jwt get jwt-like string
func Jwt(opts ...options.OptionFunc) string {
	return singleFakeData(JWT, func() interface{} {
		opt := options.BuildOptions(opts)
		i := Internet{fakerOption: *opt}
		p, err := i.jwt()
		if err != nil {
			panic(err.Error())
		}
		return p
	}, opts...).(string)
}
