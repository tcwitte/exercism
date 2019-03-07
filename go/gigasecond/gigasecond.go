// Package gigasecond offers a function for adding a gigasecond to a given time.
package gigasecond

import "time"

const testVersion = 4

const gigaSecond = 1000000000 * time.Second

// AddGigasecond returns the given time with 10^9 seconds added to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
