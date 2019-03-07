// Package pangram provides a function for computing whether a string contains
// all letters of the alphabet.
package pangram

import "unicode"

const testVersion = 1

// IsPangram returns whether the given string contains all letters of the
// alphabet, ignoring case.
func IsPangram(text string) bool {
	// Store all letters that we're interested in
	alphabet := make(map[rune]bool, 26)
	for letter := 'A'; letter <= 'Z'; letter++ {
		alphabet[letter] = false
	}

	// Mark all letters that occur
	for _, r := range text {
		letter := unicode.ToUpper(r)
		if _, ok := alphabet[letter]; ok {
			alphabet[letter] = true
		}
	}

	// Determine whether all interesting letters occur
	for _, present := range alphabet {
		if !present {
			return false
		}
	}
	return true
}
