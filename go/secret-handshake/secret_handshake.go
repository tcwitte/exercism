// Package secret contains a function to compute the secret handshake represented by a number.
package secret

const testVersion = 1

// Handshake computes the secret handshake for the given number.
func Handshake(code uint) []string {
	var result []string
	if (code & 1) != 0 {
		result = append(result, "wink")
	}
	if (code & 2) != 0 {
		result = append(result, "double blink")
	}
	if (code & 4) != 0 {
		result = append(result, "close your eyes")
	}
	if (code & 8) != 0 {
		result = append(result, "jump")
	}
	if (code&16) != 0 && len(result) != 0 {
		for i := 0; i < len(result)/2; i++ {
			result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
		}
	}

	return result
}
