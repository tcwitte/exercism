// Package clock implements a 24 hour clock with hours and minutes.
package clock

import "fmt"

const testVersion = 4

const minutesPerHour = 60
const minutesPerDay = 24 * 60

// Clock represents a 24 hour clock with hours and minutes. The time is stored as the number of minutes from the start of the day.
type Clock int

// New returns a clock initialized to the given time.
func New(hours, minutes int) Clock {
	var c Clock
	return c.Add(minutesPerHour*hours + minutes)
}

// Hours return the time in hours from 0 to 24.
func (c Clock) Hours() int {
	return int(c) / minutesPerHour
}

// Minutes returns the time in minutes after the hour.
func (c Clock) Minutes() int {
	return int(c) % minutesPerHour
}

// String returns the time as a string in 24 hour notation with 2 digits for the hours and 2 digits for the minutes. For example, 04:49.
func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.Hours(), c.Minutes())
}

// Add adds the given number of minutes to the time. The number of minutes can be negative.
func (c Clock) Add(minutes int) Clock {
	c += Clock(minutes % minutesPerDay)
	c %= minutesPerDay
	if c < 0 {
		c += minutesPerDay
	}
	return c
}
