// Package hamming provides a function for computing Hamming distances.
package hamming

import "fmt"

const testVersion = 5

// Distance returns the Hamming distance between a and b. If the strings are not of equal length, an error is returned.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings are not of equal length")
	}

	distance := 0
	bchars := []rune(b)
	for i, achar := range a {
		if achar != bchars[i] {
			distance++
		}
	}

	return distance, nil
}
