package interfaces

import "reflect"

// CustomProviderFunction used as the standard layout function for custom providers
type CustomProviderFunction func() (interface{}, error)

// TaggedFunction used as the standard layout function for tag providers in struct.
// This type also can be used for custom provider.
type TaggedFunction func(v reflect.Value) (interface{}, error)
