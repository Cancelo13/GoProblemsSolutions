package microblog

import "unicode/utf8"

func Truncate(phrase string) string {
	if utf8.RuneCountInString(phrase) <= 5 {
		return phrase
	}
	var res string = ""
	var cnt int = 5
	for _, c := range phrase {
		res += string(c)
		cnt--
		if cnt == 0 {
			break
		}
	}
	return res
}
