// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isYelling(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isYelling(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isQuestion(message string) bool {
	return strings.HasSuffix(message, "?")
}

func isYelling(message string) bool {
	upperCaseMessage := strings.ToUpper(message)
	return message == upperCaseMessage && strings.ToLower(message) != upperCaseMessage
}
