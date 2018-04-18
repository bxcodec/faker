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
	"http://%s/",
	"http://www.%s/%s",
	"http://www.%s/%s",
	"https://www.%s/%s",
	"http://%s/%s",
	"http://%s/%s",
	"http://%s/%s.html",
	"https://%s/%s.html",
	"http://%s/%s.php",
}
var internet Networker

// Constructor
func GetNetworker() Networker {
	mu.Lock()
	defer mu.Unlock()

	if internet == nil {
		internet = &Internet{}
	}
	return internet
}

// this set custom Network
func SetNetwork(net Networker) {
	internet = net
}

type Networker interface {
	Email() string
	MacAddress() string
	DomainName() string
	Url() string
	UserName() string
	Ipv4() string
	Ipv6() string
	Password() string
}

type Internet struct{}

func (internet Internet) Email() string {
	return randomString(7) + "@" + randomString(5) + "." + randomElementFromSliceString(tld)
}
func (internet Internet) MacAddress() string {
	ip := make([]byte, 6)
	for i := 0; i < 6; i++ {
		ip[i] = byte(rand.Intn(256))
	}

	return net.HardwareAddr(ip).String()
}
func (internet Internet) DomainName() string {
	return randomString(7) + "." + randomElementFromSliceString(tld)
}
func (internet Internet) Url() string {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	if countVerbs == 1 {
		return fmt.Sprintf(format, internet.DomainName())
	} else {
		return fmt.Sprintf(format, internet.DomainName(), internet.UserName())
	}
}
func (internet Internet) UserName() string {
	return randomString(7)
}
func (internet Internet) Ipv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To4().String()
}
func (internet Internet) Ipv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// get hash password
func (internet Internet) Password() string {
	return randomString(50)
}
