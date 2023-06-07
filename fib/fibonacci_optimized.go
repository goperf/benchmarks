package fib

var memo = make(map[int]int)

/*
In this optimized version, we introduce an additional helper function fibMemo that takes an additional memo parameter.
The memo parameter is a map that stores the previously calculated Fibonacci values.
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
