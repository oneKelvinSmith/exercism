package diffsquares

// SquareOfSums retuns the square of sums for the first n natural numbers.
func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares retuns the sum of squares for the first n natural numbers.
func SumOfSquares(n int) int {
	return faulhaber(n)
}

func faulhaber(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the and SquareOfSums
// SumOfSquares for the first n natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
