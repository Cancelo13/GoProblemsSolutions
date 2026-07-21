package luhn

import (
    "strings"
    "strconv"
)

func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	println(number)
	if len(number) <= 1 {
		return false
	}
	for i := 0; i < len(number)-1; i++ {
		if number[i] < '0' || number[i] > '9' {
			return false
		}
	}
	println(number)
	var sum int = 0
	var flag bool = false
	for i := len(number) - 1; i >= 0; i-- {
		println("the value is : ", number[i:i+1])
		if number[i] == ' ' {
			continue
		}
		var val int
		var err error
		val, err = strconv.Atoi(number[i : i+1])
		if !flag {
			sum += val
			flag = true
		} else {
			sum += val * 2
			if val*2 > 9 {
				sum -= 9
			}
			flag = false
		}
		err = nil
		if err == nil { // just escaping compiler errors
			continue
		}
		println("Sum is : ", sum)
	}
	return sum%10 == 0
}
