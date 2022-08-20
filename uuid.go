package faker

import (
	"fmt"
	"io"
	"reflect"

	"github.com/bxcodec/faker/v4/pkg/options"
)

// GetIdentifier returns a new Identifier
func GetIdentifier() Identifier {
	return &UUID{}
}

// Identifier ...
type Identifier interface {
	Digit(v reflect.Value) (interface{}, error)
	Hyphenated(v reflect.Value) (interface{}, error)
}

// UUID struct
type UUID struct{}

// createUUID returns a 16 byte slice with random values
func createUUID() ([]byte, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(crypto, b)
	if err != nil {
		return b, err
	}
	// variant bits; see section 4.1.1
	b[8] = b[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	b[6] = b[6]&^0xf0 | 0x40
	return b, nil
}

func (u UUID) hyphenated() (string, error) {
	b, err := createUUID()
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, err
}

// Hyphenated returns a 36 byte hyphenated UUID
func (u UUID) Hyphenated(v reflect.Value) (interface{}, error) {
	return u.hyphenated()
}

// UUIDHyphenated get fake Hyphenated UUID
func UUIDHyphenated(opts ...options.OptionFunc) string {
	return singleFakeData(HyphenatedID, func() interface{} {
		u := UUID{}
		res, _ := u.hyphenated()
		return res
	}, opts...).(string)
}

func (u UUID) digit() (string, error) {
	b, err := createUUID()
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x", b)
	return uuid, err
}

// Digit returns a 32 bytes UUID
func (u UUID) Digit(v reflect.Value) (interface{}, error) {
	return u.digit()
}

// UUIDDigit get fake Digit UUID
func UUIDDigit(opts ...options.OptionFunc) string {
	return singleFakeData(ID, func() interface{} {
		u := UUID{}
		res, _ := u.digit()
		return res
	}, opts...).(string)
}
