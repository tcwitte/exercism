// Package acronym provides a function for abbreviating a phrase to its acronym.
package acronym

import (
	"bytes"
	"unicode"
)

const testVersion = 2

// Abbreviate takes a phrase and returns its acronym. The acronym contains one
// character for each word in the phrase, but also for each capitalized part of
// a word.
func Abbreviate(phrase string) string {
	var buffer bytes.Buffer
	inWord := false      // We're currently parsing a word
	prevIsUpper := false // Previous rune is in uppercase

	for _, r := range phrase {
		if unicode.IsLetter(r) {
			if !inWord || (!prevIsUpper && unicode.IsUpper(r)) {
				inWord = true
				buffer.WriteRune(unicode.ToUpper(r))
			}

			prevIsUpper = unicode.IsUpper(r)
		} else {
			inWord = false
			prevIsUpper = false
		}
	}

	return buffer.String()
}
