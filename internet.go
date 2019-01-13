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

func (internet Internet) domainName() string {
	return randomString(7) + "." + randomElementFromSliceString(tld)
}

// DomainName generates random domain name
func (internet Internet) DomainName(v reflect.Value) (interface{}, error) {
	return internet.domainName(), nil
}

func (internet Internet) url() string {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	if countVerbs == 1 {
		return fmt.Sprintf(format, internet.domainName())
	}
	return fmt.Sprintf(format, internet.domainName(), internet.username())
}

// URL generates random URL standardised in urlFormats const
func (internet Internet) URL(v reflect.Value) (interface{}, error) {
	return internet.url(), nil
}

func (internet Internet) username() string {
	return randomString(7)
}

// UserName generates random username
func (internet Internet) UserName(v reflect.Value) (interface{}, error) {
	return internet.username(), nil
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

func (internet Internet) password() string {
	return randomString(50)
}

// Password returns a hashed password
func (internet Internet) Password(v reflect.Value) (interface{}, error) {
	return internet.password(), nil
}
