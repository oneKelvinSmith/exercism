package gigasecond

import "time"

const gigasecond = 1e9 * time.Second

// AddGigasecond - Adds a gigasecond to a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
