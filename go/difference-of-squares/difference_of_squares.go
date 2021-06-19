package diffsquares

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}

func SquareOfSum(input int) int {
	var sum int
	for i := 1; i <= input; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(input int) int {
	var sum int
	for i := 1; i <= input; i++ {
		sum += i * i
	}
	return sum
}
