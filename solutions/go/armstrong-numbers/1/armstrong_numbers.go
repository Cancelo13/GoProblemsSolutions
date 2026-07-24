package armstrongnumbers

import (
    "strconv"
    "math"
)

func IsNumber(n int) bool{
	var d int = 0
	var sum int = 0
	var N int = n
	var m int = len(strconv.Itoa(n))
	for n > 0{
		d = n % 10
		sum += int(math.Pow(float64(d), float64(m)))
		n /= 10
	}
	return sum == N
}