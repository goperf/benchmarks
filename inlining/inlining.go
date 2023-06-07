package inlining

func Add(a, b int) int {
	return a + b
}

//go:noinline
func AddNoInlined(a, b int) int {
	return a + b
}

// Querestion? How to make a function inlineable?
