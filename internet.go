package faker

import (
	"fmt"
	"math/rand"
	"net"
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
	Email() string
	MacAddress() string
	DomainName() string
	URL() string
	UserName() string
	IPv4() string
	IPv6() string
	Password() string
}

// Internet struct
type Internet struct{}

// Email generates random email id
func (internet Internet) Email() string {
	return randomString(7) + "@" + randomString(5) + "." + randomElementFromSliceString(tld)
}

// MacAddress generates random MacAddress
func (internet Internet) MacAddress() string {
	ip := make([]byte, 6)
	for i := 0; i < 6; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.HardwareAddr(ip).String()
}

// DomainName generates random domain name
func (internet Internet) DomainName() string {
	return randomString(7) + "." + randomElementFromSliceString(tld)
}

// URL generates random URL standardised in urlFormats const
func (internet Internet) URL() string {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	if countVerbs == 1 {
		return fmt.Sprintf(format, internet.DomainName())
	}
	return fmt.Sprintf(format, internet.DomainName(), internet.UserName())
}

// UserName generates random username
func (internet Internet) UserName() string {
	return randomString(7)
}

// IPv4 generates random IPv4 address
func (internet Internet) IPv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 generates random IPv6 address
func (internet Internet) IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// Password returns a hashed password
func (internet Internet) Password() string {
	return randomString(50)
}
