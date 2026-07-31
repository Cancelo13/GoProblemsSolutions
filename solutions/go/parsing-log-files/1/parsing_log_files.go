package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|FTL|ERR)\]\s`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string{
	re := regexp.MustCompile(`<[*|~|=|-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int = 0
	re := regexp.MustCompile(`\".*(?i)password\.*"`)
	for _, v := range lines{
		if re.MatchString(v){
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string{
	re := regexp.MustCompile(`\bend-of-line\d*\b`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\bUser\b\s*(\w+)`)
	for i,v := range lines{
		s := re.FindStringSubmatch(v)
		if s != nil{
			lines[i] = "[USR] " + string(s[1]) + " " + v
		}
	}
	return lines
}