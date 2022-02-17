package util

import "strings"

// IndentMultiline increases indentation of a multiline string
func IndentMultiline(multilineStr string, indendationLevel int) string {
	indentation := strings.Repeat(" ", indendationLevel)
	return indentation + strings.ReplaceAll(multilineStr, "\n", "\n"+indentation)
}
