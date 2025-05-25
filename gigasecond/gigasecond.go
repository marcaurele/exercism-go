// Package gigasecond compute the datetime interval at the giga (1'000'000'000) seconds interval.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds 1'000'000'000 seconds to the given datetime.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
