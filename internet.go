package faker

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
)

var tld = []string{"com", "com", "biz", "info", "net", "org", "ru"}
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
var mu sync.Mutex

// Constructor
func getNetworker() Networker {
	mu.Lock()
	defer mu.Unlock()

	if internet == nil {
		internet = &Internet{}
	}
	return internet
}

// this set custom Network
func SetNetwork(net Networker) {
	internet = &net
}

type Networker interface {
	Email() string
	MacAddress() string
	DomainName() string
	Url() string
	UserName() string
	Ipv4() string
	Ipv6() string
}

type Internet struct{}

func (i Internet) Email() string {
	return randomString(7) + "@" + randomString(5) + randomElementFromSliceString(tld)
}
func (i Internet) MacAddress() string {
	r := rand.New(src)
	size := 6
	ip := make([]string, size*3)
	for i := 0; i < 6; i++ {
		ip[i] = string(r.Intn(256))
	}
	mac := strings.Join(ip, ":")

	log.Println(mac)
	parseMacAddr, err := net.ParseMAC(mac)

	if err != nil {
		panic(err)
	}

	return parseMacAddr.String()
}
func (i Internet) DomainName() string {
	return randomString(7) + "." + randomElementFromSliceString(tld)
}
func (i Internet) Url() string {
	format := randomElementFromSliceString(urlFormats)
	countVerbs := strings.Count(format, "%s")
	if countVerbs == 1 {
		return fmt.Sprintf(format, i.DomainName())
	} else {
		return fmt.Sprintf(format, i.DomainName(), i.UserName())
	}
}
func (i Internet) UserName() string {
	return randomString(7)
}
func (i Internet) Ipv4() string {
	r := rand.New(src)
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(r.Intn(256))
	}
	return net.IP(ip).To4().String()
}
func (i Internet) Ipv6() string {
	r := rand.New(src)
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(r.Intn(256))
	}
	return net.IP(ip).To16().String()
}
