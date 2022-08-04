package Utilities

import (
	"regexp"
	"strings"
)

func TryFindScriptMatch(expressionPtr *regexp.Regexp, searchBodyPtr *string, searchQuery string) *string {

	matches := expressionPtr.FindAllString(*searchBodyPtr, 15)
	for i := range matches {
		if strings.Contains(matches[i], searchQuery) {
			return &matches[i]
		}
	}
	return nil
}
