package faker

import "reflect"

type options struct {
	ignoreFields      map[string]struct{}
	fieldProviders    map[string]CustomProviderFunction
	typeSeen          map[reflect.Type]int
	recursionMaxDepth int
}

func (o *options) rememberType(t reflect.Type) {
	o.typeSeen[t]++
}

func (o *options) forgetType(t reflect.Type) {
	o.typeSeen[t]--
}

func (o *options) recursionOutOfLimit(t reflect.Type) bool {
	return o.typeSeen[t] > o.recursionMaxDepth
}

func buildOptions(optFuncs []OptionFunc) *options {
	ops := &options{}
	ops.typeSeen = make(map[reflect.Type]int, 1)
	ops.recursionMaxDepth = 1
	for _, optFunc := range optFuncs {
		optFunc(ops)
	}
	return ops
}

type OptionFunc func(oo *options)

func WithFieldsToIgnore(fieldNames ...string) OptionFunc {
	return func(oo *options) {
		if oo.ignoreFields == nil {
			oo.ignoreFields = make(map[string]struct{}, len(fieldNames))
		}
		for _, f := range fieldNames {
			oo.ignoreFields[f] = struct{}{}
		}
	}
}

func WithCustomFieldProvider(fieldName string, provider CustomProviderFunction) OptionFunc {
	return func(oo *options) {
		if oo.fieldProviders == nil {
			oo.fieldProviders = make(map[string]CustomProviderFunction, 1)
		}
		oo.fieldProviders[fieldName] = provider
	}
}

func WithRecursionMaxDepth(depth int) OptionFunc {
	return func(oo *options) {
		if depth >= 0 {
			oo.recursionMaxDepth = depth
		}
	}
}
