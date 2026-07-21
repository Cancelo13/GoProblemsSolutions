package collatzconjecture

import "errors"

func CollatzConjecture(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("Error")
	}
	var counter int = 0
	for number > 1 {
		if number%2 == 0 {
			number /= 2
			counter++
		} else {
			number *= 3
			number++
			counter++
		}
	}
	return counter, nil
}
