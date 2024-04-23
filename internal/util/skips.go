package util

import "strings"

func IsLineComment(ln string) bool {
	return strings.HasPrefix(ln, "//")
}

func IsEmptyLine(ln string) bool {
	return len(strings.TrimSpace(ln)) < 1
}
