package bmark

//go:noinline
func EscapeAnalysis() []int {
	var data = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return data
}

//go:noinline
func EscapeAnalysisOptimized() [10]int {
	var data = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return data
}
