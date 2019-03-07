// Package bob implements a lackadaisical teenager whose responses are very
// limited.
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

// Hey responds to the given text in a lackadaisical way.
func Hey(text string) (response string) {
	text = strings.TrimFunc(text, unicode.IsSpace)

	switch {
	case len(text) == 0:
		response = "Fine. Be that way!"
	case strings.ToUpper(text) == text && strings.ToLower(text) != text:
		// Uppercase text that contains letters. May include yelled questions.
		response = "Whoa, chill out!"
	case text[len(text)-1] == '?':
		response = "Sure."
	default:
		response = "Whatever."
	}

	return
}
