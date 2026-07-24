package scrabblescore

import "strings"

func Score(word string) int {
	var c string = ""
	var sum int = 0
	for i := 0; i < len(word); i++ {
		// it can easily solved by a map like this : sum += mp[c]
		c = strings.ToUpper(string(word[i]))
		switch c {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			sum++
		case "D", "G":
			sum += 2
		case "B", "C", "M", "P":
			sum += 3
		case "F", "H", "V", "W", "Y":
			sum += 4
		case "K":
			sum += 5
		case "J", "X":
			sum += 8
		case "Q", "Z":
			sum += 10
		default:
			sum = sum
		}
	}
	return sum
}