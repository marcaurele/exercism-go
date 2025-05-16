package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var reIsValidLine = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

func IsValidLine(text string) bool {
	return reIsValidLine.MatchString(text)
}

var reSplitLogLine = regexp.MustCompile(`<(\*|~|=|-)*>`)

func SplitLogLine(text string) []string {
	return reSplitLogLine.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var reCountQuotedPasswords = regexp.MustCompile(`(?i)".*password.*"`)
	var total int
	for i := 0; i < len(lines); i++ {
		if reCountQuotedPasswords.MatchString(lines[i]) {
			total++
		}
	}
	return total
}

var reRemoveEndOfLineText = regexp.MustCompile(`end-of-line\d*`)

func RemoveEndOfLineText(text string) string {
	return reRemoveEndOfLineText.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var reTagWithUserName = regexp.MustCompile(`User\s+(\w+)`)
	result := []string{}
	for _, line := range lines {
		if match := reTagWithUserName.FindStringSubmatch(line); match != nil {
			line = fmt.Sprintf("[USR] %s %s", match[1], line)
		}
		result = append(result, line)
	}
	return result
}
