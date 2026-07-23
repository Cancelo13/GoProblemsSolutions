package anagram

import (
    "strings"
    "slices"
)

func Detect(original string, words []string) []string {
	original = strings.ToLower(original)
	runes := []rune(original)
	slices.Sort(runes)
	sortedOriginal := string(runes)
	var res []string
	var newWords []string = []string{}
	newWords = append(newWords, words...)
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		if words[i] == original {
			continue
		} else {
			var flg bool = true
			if len(words[i]) != len(original) {
				continue
			}
			// Sort words[i], original then compare character by character
			runes = []rune(words[i])
			slices.Sort(runes)
			words[i] = string(runes)
			for j := 0; j < len(words[i]); j++ {
				var c string = string(words[i][j])
				if c != string(sortedOriginal[j]) {
					flg = false
					break
				}
			}
			if flg {
				res = append(res, newWords[i])
			}
		}
	}
	return res
}
