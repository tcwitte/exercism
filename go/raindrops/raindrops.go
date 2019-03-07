// Package raindrops provides a number to string converter that takes into
// account the number's factors.
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

// Convert takes an integer and returns its string-representation, except if the
// number is divisible by 3, 5 or 7. Then it prints Pling, Plang and/or Plong
// respectively.
func Convert(number int) string {
	var buffer bytes.Buffer

	if number%3 == 0 {
		buffer.WriteString("Pling")
	}

	if number%5 == 0 {
		buffer.WriteString("Plang")
	}

	if number%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() != 0 {
		return buffer.String()
	}
	return strconv.FormatInt(int64(number), 10)
}
