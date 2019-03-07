// Package diffsquares provides functions for computing the square of sums and the sum of squares of
package diffsquares

const testVersion = 1

// SquareOfSums returns the square of the sum of the first n natural numbers.
func SquareOfSums(n int) int {
	result := 0
	for i := n; i > 0; i-- {
		result += i
	}
	return result * result
}

// SumOfSquares returns the sum of squares of the first n natural numbers.
func SumOfSquares(n int) int {
	if n <= 0 {
		return 0
	}
	return n*n + SumOfSquares(n-1)
}

// Difference returns the difference between the sum of squares and the square of sums. It is always positive.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
