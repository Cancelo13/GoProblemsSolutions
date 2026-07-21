package differenceofsquares

func SquareOfSum(number int) int {
	var sum int = number * (number + 1) / 2
	return sum * sum
}

func SumOfSquares(number int) int{
	return number * (number + 1) * (2 * number + 1) / 6
}

func Difference(number int) int{
	return SquareOfSum(number) - SumOfSquares(number)
}