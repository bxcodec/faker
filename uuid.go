package faker

import (
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
)

var identifier Identifier

// GetIdentifier returns a new Identifier
func GetIdentifier() Identifier {
	mu.Lock()
	defer mu.Unlock()

	if identifier == nil {
		identifier = &UUID{}
	}
	return identifier
}

// Identifier ...
type Identifier interface {
	Digit(v reflect.Value) (interface{}, error)
	Hyphenated(v reflect.Value) (interface{}, error)
}

// UUID struct
type UUID struct{}

// createUUID returns a 16 byte slice with random values
func createUUID() []byte {
	b := make([]byte, 16)
	io.ReadFull(rand.Reader, b)
	// variant bits; see section 4.1.1
	b[8] = b[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	b[6] = b[6]&^0xf0 | 0x40
	return b
}

func (u UUID) hyphenated() string {
	b := createUUID()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// Hyphenated returns a 36 byte hyphenated UUID
func (i UUID) Hyphenated(v reflect.Value) (interface{}, error) {
	return i.hyphenated(), nil
}

func (u UUID) digit() string {
	b := createUUID()
	uuid := fmt.Sprintf("%x", b)
	return uuid
}

// Digit returns a 32 bytes UUID
func (i UUID) Digit(v reflect.Value) (interface{}, error) {

	return i.digit(), nil
}
