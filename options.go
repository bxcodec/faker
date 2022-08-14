package faker

type options struct {
	ignoreFields   map[string]struct{}
	fieldProviders map[string]CustomProviderFunction
}

func buildOptions(optFuncs []OptionFunc) *options {
	ops := &options{}
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
