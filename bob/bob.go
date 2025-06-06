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
	spaceTrimmed := strings.TrimSpace(remark)
	switch {
	case spaceTrimmed == "":
		return "Fine. Be that way!"
	case strings.HasSuffix(spaceTrimmed, "?") && spaceTrimmed == strings.ToUpper(spaceTrimmed) && strings.ContainsAny(strings.ToUpper(spaceTrimmed), "ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		return "Calm down, I know what I'm doing!"
	case spaceTrimmed == strings.ToUpper(spaceTrimmed) && strings.ContainsAny(spaceTrimmed, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"):
		return "Whoa, chill out!"
	case strings.HasSuffix(spaceTrimmed, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
