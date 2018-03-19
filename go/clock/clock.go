package clock

import "fmt"

const (
	minutesInHour = 60
	minutesInDay  = 24 * 60
)

// New constructs a new Clock.
func New(hour, minute int) Clock {
	hour, minute = normalise(hour*minutesInHour + minute)

	return Clock{hour: hour, minute: minute}
}

func normalise(minutes int) (hour, minute int) {
	minutes = (minutesInDay + minutes) % minutesInDay

	if minutes < 0 {
		return normalise(minutes)
	}

	hour = minutes / minutesInHour
	minute = minutes % minutesInHour

	return
}

// Clock represents a timekeeping construct.
type Clock struct {
	hour   int
	minute int
}

// Add moves the clock time forward by the given number of minutes.
func (c Clock) Add(minutes int) Clock {
	clockMinutes := c.hour*minutesInHour + c.minute
	hour, minute := normalise(clockMinutes + minutes)

	return Clock{hour: hour, minute: minute}
}

// Subtract moves the clock time back by the given number of minutes.
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
