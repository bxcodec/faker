package faker

import (
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
)

var uuid Identifier

// GetUUID returns a new Uuid interface of Uuid
func GetUUID() Identifier {
	mu.Lock()
	defer mu.Unlock()

	if uuid == nil {
		uuid = &UUID{}
	}
	return uuid
}

// SetUUID sets custom UUID
func SetUUID(identifier UUID) {
	uuid = identifier
}

// Identifier ...
type Identifier interface {
	Digit(v reflect.Value) error
	Hyphenated(v reflect.Value) error
}

// UUID struct
type UUID struct{}

func createUUID() ([]byte, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}
	// variant bits; see section 4.1.1
	b[8] = b[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	b[6] = b[6]&^0xf0 | 0x40
	return b, nil
}

// Hyphenated returns a 36 bytes UUID with hyphens
func (i UUID) Hyphenated(v reflect.Value) error {
	b, err := createUUID()
	if err != nil {
		return err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	v.Set(reflect.ValueOf(uuid))
	return nil
}

// Digit returns a 32 bytes UUID with chars
func (i UUID) Digit(v reflect.Value) error {
	b, err := createUUID()
	if err != nil {
		return err
	}
	uuid := fmt.Sprintf("%x", b)
	v.Set(reflect.ValueOf(uuid))
	return nil
}
