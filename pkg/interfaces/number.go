package interfaces

// RandomIntegerBoundary is the struct for random integer boundaries
type RandomIntegerBoundary struct {
	Start int
	End   int
}

// RandomFloatBoundary is the struct for random float boundaries
type RandomFloatBoundary struct {
	Start float64
	End   float64
}

var (
	DefaultIntBoundary   = RandomIntegerBoundary{Start: 0, End: 100}
	DefaultFloatBoundary = RandomFloatBoundary{Start: 0, End: 100}
)
