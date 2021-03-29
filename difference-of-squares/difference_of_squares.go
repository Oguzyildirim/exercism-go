//  Package diffsquares calculates the square of the sum and the sum of the squares
package diffsquares

// SquareOfSum calculates square of the sum
func SquareOfSum(n int) int {
	var result int
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	result = sum * sum
	return result
}

// SumOfSquares calculates sum of the squares
func SumOfSquares(n int) int {
	var result int
	for i := 0; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference calculates difference between the square of the sum and the sum of the squares
func Difference(n int) int {
	result := SquareOfSum(n) - SumOfSquares(n)
	return result
}
