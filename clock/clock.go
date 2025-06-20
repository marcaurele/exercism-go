package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {
	totalMinutes := h*60 + m
	adjust := 0
	if totalMinutes%60 < 0 {
		adjust = 1
	}

	return Clock{
		Hour:   ((totalMinutes/60)%24 + 24 - adjust) % 24,
		Minute: ((totalMinutes % 60) + 60) % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.Hour, c.Minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
