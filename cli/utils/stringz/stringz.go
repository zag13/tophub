package stringz

import "unicode/utf8"

func TruncateString(s string, maxLength int) string {
	if utf8.RuneCountInString(s) <= maxLength {
		return s
	}

	truncated := []rune(s)[:maxLength-3]
	return string(truncated) + "..."
}
