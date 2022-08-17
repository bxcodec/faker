package options

import (
	"fmt"
	"reflect"

	fakerErrors "github.com/bxcodec/faker/v4/errors"
	"github.com/bxcodec/faker/v4/interfaces"
)

type Options struct {
	IgnoreFields             map[string]struct{}
	FieldProviders           map[string]interfaces.CustomProviderFunction
	MaxDepthOption           *MaxDepthOption
	IgnoreInterface          bool
	StringLanguage           *interfaces.LangRuneBoundary
	GenerateUniqueValues     bool
	RandomStringLength       int
	RandomMaxSliceSize       int
	RandomMinSliceSize       int
	MaxGenerateStringRetries int
}

type MaxDepthOption struct {
	typeSeen          map[reflect.Type]int
	recursionMaxDepth int
}

func (o *MaxDepthOption) RememberType(t reflect.Type) {
	o.typeSeen[t]++
}

func (o *MaxDepthOption) ForgetType(t reflect.Type) {
	o.typeSeen[t]--
}

func (o *MaxDepthOption) RecursionOutOfLimit(t reflect.Type) bool {
	return o.typeSeen[t] > o.recursionMaxDepth
}

func BuildOptions(optFuncs []OptionFunc) *Options {
	ops := DefaultOption()

	for _, optFunc := range optFuncs {
		optFunc(ops)
	}

	return ops
}

func DefaultOption() *Options {
	ops := &Options{
		MaxDepthOption: &MaxDepthOption{},
	}
	ops.MaxDepthOption.typeSeen = make(map[reflect.Type]int, 1)
	ops.MaxDepthOption.recursionMaxDepth = 1 // default
	ops.StringLanguage = &interfaces.LangENG
	ops.RandomStringLength = 25            //default
	ops.RandomMaxSliceSize = 100           //default
	ops.RandomMinSliceSize = 0             // default
	ops.MaxGenerateStringRetries = 1000000 //default
	return ops
}

type OptionFunc func(oo *Options)

func WithFieldsToIgnore(fieldNames ...string) OptionFunc {
	return func(oo *Options) {
		if oo.IgnoreFields == nil {
			oo.IgnoreFields = make(map[string]struct{}, len(fieldNames))
		}
		for _, f := range fieldNames {
			oo.IgnoreFields[f] = struct{}{}
		}
	}
}

func WithCustomFieldProvider(fieldName string, provider interfaces.CustomProviderFunction) OptionFunc {
	return func(oo *Options) {
		if oo.FieldProviders == nil {
			oo.FieldProviders = make(map[string]interfaces.CustomProviderFunction, 1)
		}
		oo.FieldProviders[fieldName] = provider
	}
}

func WithRecursionMaxDepth(depth uint) OptionFunc {
	return func(oo *Options) {
		if oo.MaxDepthOption == nil {
			oo.MaxDepthOption = &MaxDepthOption{
				recursionMaxDepth: 1, // default
				typeSeen:          make(map[reflect.Type]int, 1),
			}
		}
		if depth >= 0 {
			oo.MaxDepthOption.recursionMaxDepth = int(depth)
		}
	}
}

// WithIgnoreInterface allows to set a flag to ignore found interface{}s.
func WithIgnoreInterface(value bool) OptionFunc {
	return func(oo *Options) {
		oo.IgnoreInterface = value
	}
}

// WithStringLanguage sets language of random string generation (LangENG, LangCHI, LangRUS, LangJPN, LangKOR, EmotEMJ)
func WithStringLanguage(l interfaces.LangRuneBoundary) OptionFunc {
	return func(oo *Options) {
		oo.StringLanguage = &l
	}
}

// WithGenerateUniqueValues allows to set the single fake data generator functions to generate unique data.
func WithGenerateUniqueValues(unique bool) OptionFunc {
	return func(oo *Options) {
		oo.GenerateUniqueValues = unique
	}
}

// WithRandomStringLength sets a length for random string generation
func WithRandomStringLength(size uint) OptionFunc {
	if size < 0 {
		err := fmt.Errorf(fakerErrors.ErrSmallerThanZero, size)
		panic(err)
	}
	return func(oo *Options) {
		oo.RandomStringLength = int(size)
	}
}

// WithRandomMapAndSliceMaxSize sets the max size for maps and slices for random generation.
func WithRandomMapAndSliceMaxSize(size uint) OptionFunc {
	if size < 1 {
		err := fmt.Errorf(fakerErrors.ErrSmallerThanOne, size)
		panic(err)
	}
	return func(oo *Options) {
		oo.RandomMaxSliceSize = int(size)
	}
}

// WithRandomMapAndSliceMinSize sets the min size for maps and slices for random generation.
func WithRandomMapAndSliceMinSize(size uint) OptionFunc {
	if size < 0 {
		err := fmt.Errorf(fakerErrors.ErrSmallerThanZero, size)
		panic(err)
	}
	return func(oo *Options) {
		oo.RandomMinSliceSize = int(size)
	}
}

// MaxGenerateStringRetries set how much tries for generating random string
func WithMaxGenerateStringRetries(retries uint) OptionFunc {
	return func(oo *Options) {
		oo.MaxGenerateStringRetries = int(retries)
	}
}
