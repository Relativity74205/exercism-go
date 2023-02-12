package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(-|~|=|\*)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	count := 0
	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s*(\w*)`)
	var newLines []string
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 0 {
			newLines = append(newLines, "[USR] "+match[1]+" "+line)
		} else {
			newLines = append(newLines, line)
		}
	}
	return lines
}
