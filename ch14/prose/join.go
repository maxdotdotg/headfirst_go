package prose

import "strings"

// JoinWithCommas takes a slice of strings and returns a comma-delimited string
func JoinWithCommas(phrases []string) string {
	// join all of the phrases except the last
	result := strings.Join(phrases[:len(phrases)-1], ", ")

	result += ", and "

	// add the last phrase
	result += phrases[len(phrases)-1]
	return result
}
