package options

import (
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"

	fakerErrors "github.com/bxcodec/faker/v4/pkg/errors"
	"github.com/bxcodec/faker/v4/pkg/interfaces"
)

var (
	// global settings, read/write must be concurrent safe
	generateUniqueValues atomic.Value
	ignoreInterface      atomic.Value
	randomStringLen      int32 = 25
	lang                 unsafe.Pointer
	randomMaxSize        int32 = 100
	randomMinSize        int32 = 0
	iBoundary            unsafe.Pointer
)

func init() {
	generateUniqueValues.Store(false)
	ignoreInterface.Store(false)
	lang = unsafe.Pointer(&interfaces.LangENG)
	iBoundary = unsafe.Pointer(&interfaces.DefaultIntBoundary)
}

// Options represent all available option for faker.
type Options struct {
	// IgnoreFields used for ignoring a field when generating the fake data
	IgnoreFields map[string]struct{}
	// FieldProviders used for storing the custom provider function
	FieldProviders map[string]interfaces.CustomProviderFunction
	// MaxDepthOption used for configuring the max depth of nested struct for faker
	MaxDepthOption *MaxDepthOption
	// IgnoreInterface used for ignoring any interface field
	IgnoreInterface bool
	// StringLanguage used for setting the language for any string in faker
	StringLanguage *interfaces.LangRuneBoundary
	// GenerateUniqueValues to ensure the generated data is unique
	GenerateUniqueValues bool
	// RandomStringLength to ensure the generated string is expected as we want
	RandomStringLength int
	// RandomMaxSliceSize used for setting the maximum of slice size, or map size that will be generated
	RandomMaxSliceSize int
	// RandomMinSliceSize used for setting the minimum of slize, array, map size that will be generated
	RandomMinSliceSize int
	// MaxGenerateStringRetries set how much tries for generating random string
	MaxGenerateStringRetries int
	// SetSliceMapNilIfLenZero allows to set nil for the slice and maps, if size is 0.
	SetSliceMapNilIfLenZero bool
	// SetSliceMapRandomToZero sets random integer generation to zero for slice and maps
	SetSliceMapRandomToZero bool
	// RandomIntegerBoundary sets boundary random integer value generation. Boundaries can not exceed integer(4 byte...)
	RandomIntegerBoundary *interfaces.RandomIntegerBoundary
	// RandomFloatBoundary sets the boundary for random float value generation. Boundaries should comply with float values constraints (IEEE 754)
	RandomFloatBoundary *interfaces.RandomFloatBoundary
}

// MaxDepthOption used for configuring the max depth of nested struct for faker
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

// BuildOptions build all option functions into one option
func BuildOptions(optFuncs []OptionFunc) *Options {
	ops := DefaultOption()

	for _, optFunc := range optFuncs {
		optFunc(ops)
	}

	return ops
}

// DefaultOption build the default option
func DefaultOption() *Options {
	ops := &Options{}
	ops.MaxDepthOption = &MaxDepthOption{
		typeSeen:          make(map[reflect.Type]int, 1),
		recursionMaxDepth: 1,
	}
	ops.GenerateUniqueValues = generateUniqueValues.Load().(bool)
	ops.IgnoreInterface = ignoreInterface.Load().(bool)
	ops.StringLanguage = (*interfaces.LangRuneBoundary)(atomic.LoadPointer(&lang))
	ops.RandomStringLength = int(atomic.LoadInt32(&randomStringLen))
	ops.RandomMaxSliceSize = int(atomic.LoadInt32(&randomMaxSize))
	ops.RandomMinSliceSize = int(atomic.LoadInt32(&randomMinSize))
	ops.MaxGenerateStringRetries = 1000000 //default
	ops.RandomIntegerBoundary = (*interfaces.RandomIntegerBoundary)(atomic.LoadPointer(&iBoundary))
	ops.RandomFloatBoundary = &interfaces.DefaultFloatBoundary
	return ops
}

// OptionFunc define the options contract
type OptionFunc func(oo *Options)

// WithFieldsToIgnore used for ignoring a field when generating the fake data
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

// WithCustomFieldProvider used for storing the custom provider function
func WithCustomFieldProvider(fieldName string, provider interfaces.CustomProviderFunction) OptionFunc {
	return func(oo *Options) {
		if oo.FieldProviders == nil {
			oo.FieldProviders = make(map[string]interfaces.CustomProviderFunction, 1)
		}
		oo.FieldProviders[fieldName] = provider
	}
}

// WithRecursionMaxDepth used for configuring the max depth of recursion struct for faker
func WithRecursionMaxDepth(depth uint) OptionFunc {
	return func(oo *Options) {
		if oo.MaxDepthOption == nil {
			oo.MaxDepthOption = &MaxDepthOption{
				recursionMaxDepth: 1, // default
				typeSeen:          make(map[reflect.Type]int, 1),
			}
		}
		oo.MaxDepthOption.recursionMaxDepth = int(depth)
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
	return func(oo *Options) {
		oo.RandomMinSliceSize = int(size)
	}
}

// WithMaxGenerateStringRetries set how much tries for generating random string
func WithMaxGenerateStringRetries(retries uint) OptionFunc {
	return func(oo *Options) {
		oo.MaxGenerateStringRetries = int(retries)
	}
}

// WithNilIfLenIsZero allows to set nil for the slice and maps, if size is 0.
func WithNilIfLenIsZero(setNil bool) OptionFunc {
	return func(oo *Options) {
		oo.SetSliceMapNilIfLenZero = setNil
	}
}

// WithSliceMapRandomToZero Sets random integer generation to zero for slice and maps
func WithSliceMapRandomToZero(setNumberToZero bool) OptionFunc {
	return func(oo *Options) {
		oo.SetSliceMapRandomToZero = setNumberToZero
	}
}

// WithRandomIntegerBoundaries sets boundary random integer value generation. Boundaries can not exceed integer(4 byte...)
func WithRandomIntegerBoundaries(boundary interfaces.RandomIntegerBoundary) OptionFunc {
	if boundary.Start > boundary.End {
		err := errors.New(fakerErrors.ErrStartValueBiggerThanEnd)
		panic(err)
	}
	return func(oo *Options) {
		oo.RandomIntegerBoundary = &boundary
	}
}

// WithRandomFloatBoundaries sets the boundary for random float value generation. Boundaries should comply with float values constraints (IEEE 754)
func WithRandomFloatBoundaries(boundary interfaces.RandomFloatBoundary) OptionFunc {
	if boundary.Start > boundary.End {
		err := errors.New(fakerErrors.ErrStartValueBiggerThanEnd)
		panic(err)
	}
	return func(oo *Options) {
		oo.RandomFloatBoundary = &boundary
	}
}

// SetGenerateUniqueValues allows to set the single fake data generator functions to generate unique data.
func SetGenerateUniqueValues(unique bool) {
	generateUniqueValues.Store(unique)
}

// SetIgnoreInterface allows to set a flag to ignore found interface{}s.
func SetIgnoreInterface(ignore bool) {
	ignoreInterface.Store(ignore)
}

// SetRandomStringLength sets a length for random string generation
func SetRandomStringLength(size int) error {
	if size < 0 {
		return fmt.Errorf(fakerErrors.ErrSmallerThanZero, size)
	}
	atomic.StoreInt32(&randomStringLen, int32(size))
	return nil
}

// SetStringLang sets language of random string generation (LangENG, LangCHI, LangRUS, LangJPN, LangKOR, EmotEMJ)
func SetStringLang(l interfaces.LangRuneBoundary) {
	atomic.StorePointer(&lang, unsafe.Pointer(&l))
}

// SetRandomMapAndSliceSize sets the size for maps and slices for random generation.
// deprecates, currently left for old version usage
func SetRandomMapAndSliceSize(size int) error {
	return SetRandomMapAndSliceMaxSize(size)
}

// SetRandomMapAndSliceMaxSize sets the max size for maps and slices for random generation.
func SetRandomMapAndSliceMaxSize(size int) error {
	if size < 1 {
		return fmt.Errorf(fakerErrors.ErrSmallerThanOne, size)
	}
	atomic.StoreInt32(&randomMaxSize, int32(size))
	return nil
}

// SetRandomMapAndSliceMinSize sets the min size for maps and slices for random generation.
func SetRandomMapAndSliceMinSize(size int) error {
	if size < 0 {
		return fmt.Errorf(fakerErrors.ErrSmallerThanZero, size)
	}
	atomic.StoreInt32(&randomMinSize, int32(size))
	return nil
}

// SetRandomNumberBoundaries sets boundary for random number generation
func SetRandomNumberBoundaries(start, end int) error {
	if start > end {
		return errors.New(fakerErrors.ErrStartValueBiggerThanEnd)
	}
	ptr := &interfaces.RandomIntegerBoundary{Start: start, End: end}
	atomic.StorePointer(&iBoundary, unsafe.Pointer(ptr))
	return nil
}
