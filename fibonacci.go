package bmark

/*
Fib returns the nth number in the Fibonacci series.
Ex:
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89 ...
*/
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
