package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var sum int = 0
	var counter = 10
	for i := 0; i < len(isbn); i++ {
		if isbn[i] == 'X' && (i == 0 || i == 9) {
			sum += 10 * counter
		} else {
			sum += int(isbn[i]-'0') * counter
		}
		counter--
	}
	return sum%11 == 0
}
