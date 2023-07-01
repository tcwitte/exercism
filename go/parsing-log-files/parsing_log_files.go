package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLine       *regexp.Regexp = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	lineSeparator   *regexp.Regexp = regexp.MustCompile(`<[~\*=\-]*>`)
	passwordPattern *regexp.Regexp = regexp.MustCompile(`".*[pP][aA][sS][sS][wW][oO][rR][dD].*"`)
	endOfLineText   *regexp.Regexp = regexp.MustCompile(`end\-of\-line[\d]+`)
	userPattern     *regexp.Regexp = regexp.MustCompile(`User\s+(\w+)`)
)

func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return lineSeparator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, l := range lines {
		if passwordPattern.MatchString(l) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return endOfLineText.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	copy(result, lines)
	for i, l := range lines {
		matches := userPattern.FindStringSubmatch(l)
		if matches != nil {
			result[i] = fmt.Sprintf("[USR] %s %s", matches[1], l)
		}
	}

	return result
}
