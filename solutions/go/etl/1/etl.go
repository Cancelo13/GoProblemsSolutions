package etl

import "strings"

func Transform(in map[int][]string) map[string]int{
	var mp map[string]int = map[string]int{}
	for i,j := range in {
		for k := 0 ; k < len(j) ; k++{
			mp[strings.ToLower(string(j[k]))] = i
		}
	}
	return mp
}
