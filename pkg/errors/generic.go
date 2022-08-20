package errors

// Generic Error Messages for tags
// 		ErrUnsupportedKindPtr: Error when get fake from ptr
// 		ErrUnsupportedKind: Error on passing unsupported kind
// 		ErrValueNotPtr: Error when value is not pointer
// 		ErrTagNotSupported: Error when tag is not supported
// 		ErrTagAlreadyExists: Error when tag exists and call AddProvider
// 		ErrTagDoesNotExist: Error when tag does not exist and call RemoveProvider
// 		ErrMoreArguments: Error on passing more arguments
// 		ErrNotSupportedPointer: Error when passing unsupported pointer
var (
	ErrUnsupportedKindPtr  = "Unsupported kind: %s Change Without using * (pointer) in Field of %s"
	ErrUnsupportedKind     = "Unsupported kind: %s"
	ErrValueNotPtr         = "Not a pointer value"
	ErrTagNotSupported     = "Tag unsupported: %s"
	ErrTagAlreadyExists    = "Tag exists"
	ErrTagDoesNotExist     = "Tag does not exist"
	ErrMoreArguments       = "Passed more arguments than is possible : (%d)"
	ErrNotSupportedPointer = "Use sample:=new(%s)\n faker.FakeData(sample) instead"
	ErrSmallerThanZero     = "Size:%d is smaller than zero."
	ErrSmallerThanOne      = "Size:%d is smaller than one."
	ErrUniqueFailure       = "Failed to generate a unique value for field \"%s\""

	ErrStartValueBiggerThanEnd = "Start value can not be bigger than end value."
	ErrWrongFormattedTag       = "Tag \"%s\" is not written properly"
	ErrUnknownType             = "Unknown Type"
	ErrNotSupportedTypeForTag  = "Type is not supported by tag."
	ErrUnsupportedTagArguments = "Tag arguments are not compatible with field type."
	ErrDuplicateSeparator      = "Duplicate separator for tag arguments."
	ErrNotEnoughTagArguments   = "Not enough arguments for tag."
	ErrUnsupportedNumberType   = "Unsupported Number type."
)
