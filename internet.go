package faker

import (
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"strings"
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
var internet Networker

// GetNetworker returns a new Networker interface of Internet
func GetNetworker() Networker {
	mu.Lock()
	defer mu.Unlock()

	if internet == nil {
		internet = &Internet{}
	}
	return internet
}

// SetNetwork sets custom Network
func SetNetwork(net Networker) {
	internet = net
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
}

// Internet struct
type Internet struct{}

func (internet Internet) email() string {
	return randomString(7) + "@" + randomString(5) + "." + randomElementFromSliceString(tld)
}

// Email generates random email id
func (internet Internet) Email(v reflect.Value) (interface{}, error) {
	return internet.email(), nil
}

// Email get email randomly in string
func Email() string {
	return singleFakeData(EmailTag, func() interface{} {
		i := Internet{}
		return i.email()
	}).(string)
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
func MacAddress() string {
	return singleFakeData(MacAddressTag, func() interface{} {
		i := Internet{}
		return i.macAddress()
	}).(string)
}

func (internet Internet) domainName() string {
	return randomString(7) + "." + randomElementFromSliceString(tld)
}

// DomainName generates random domain name
func (internet Internet) DomainName(v reflect.Value) (interface{}, error) {
	return internet.domainName(), nil
}

// DomainName get email domain name in string
func DomainName() string {
	return singleFakeData(DomainNameTag, func() interface{} {
		i := Internet{}
		return i.domainName()
	}).(string)
}

func (internet Internet) url() string {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	if countVerbs == 1 {
		return fmt.Sprintf(format, internet.domainName())
	}
	return fmt.Sprintf(format, internet.domainName(), internet.username())
}

// URL generates random URL standardized in urlFormats const
func (internet Internet) URL(v reflect.Value) (interface{}, error) {
	return internet.url(), nil
}

// URL get Url randomly in string
func URL() string {
	return singleFakeData(URLTag, func() interface{} {
		i := Internet{}
		return i.url()
	}).(string)
}

func (internet Internet) username() string {
	return randomString(7)
}

// UserName generates random username
func (internet Internet) UserName(v reflect.Value) (interface{}, error) {
	return internet.username(), nil
}

// Username get username randomly in string
func Username() string {
	return singleFakeData(UserNameTag, func() interface{} {
		i := Internet{}
		return i.username()
	}).(string)
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
func IPv4() string {
	return singleFakeData(IPV4Tag, func() interface{} {
		i := Internet{}
		return i.ipv4()
	}).(string)
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
func IPv6() string {
	return singleFakeData(IPV6Tag, func() interface{} {
		i := Internet{}
		return i.ipv6()
	}).(string)
}

func (internet Internet) password() string {
	return randomString(50)
}

// Password returns a hashed password
func (internet Internet) Password(v reflect.Value) (interface{}, error) {
	return internet.password(), nil
}

// Password get password randomly in string
func Password() string {
	return singleFakeData(PASSWORD, func() interface{} {
		i := Internet{}
		return i.password()
	}).(string)
}
