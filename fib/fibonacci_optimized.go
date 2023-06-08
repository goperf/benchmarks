package fib

var memo = make(map[int]int)

/*
FibOptimized uses a technique called memoization to improve the performance.
The memo variable is a map that stores the previously calculated Fibonacci values.
*/
func FibOptimized(n int) int {
	if n < 2 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = FibOptimized(n-1) + FibOptimized(n-2)
	return memo[n]
}
