package gigasecond

import "time"

// AddGigasecond - Adds a gigasecond to a given time
func AddGigasecond(t time.Time) time.Time {
	gigasecond := 1000000000 * time.Second

	return t.Add(gigasecond)
}
