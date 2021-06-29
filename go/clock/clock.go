package clock

import "fmt"

type Clock struct {
	hh int
	mm int
}

func New(hour, minute int) Clock {
	hour = getHours(hour)
	hoursLeft, minutesLeft := getTimeByMin(minute)
	minute = minutesLeft
	if hour+hoursLeft >= 24 {
		hour += hoursLeft - 24
	} else {
		hour += hoursLeft
	}
	if minute < 0 {
		hour--
		minute += 60
	}
	hour = hour % 24
	if hour < 0 {
		hour += 24
	}
	return Clock{hour, minute}
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hh, c.mm)
}
func (c Clock) Add(minutes int) Clock {
	hours, minutesLeft := getTimeByMin(minutes)
	if c.mm+minutesLeft >= 60 {
		c.mm += minutesLeft - 60
		c.hh++
	} else {
		c.mm += minutesLeft
	}
	if c.hh+hours > 23 {
		c.hh += hours - 24
	} else {
		c.hh += hours
	}
	return c
}
func (c Clock) Subtract(minutes int) Clock {
	hours, minutesLeft := getTimeByMin(minutes)
	if c.mm-minutesLeft < 0 {
		c.mm -= (minutesLeft - 60)
		c.hh--
	} else {
		c.mm -= minutesLeft
	}
	if c.hh-hours < 0 {
		c.hh -= (hours - 24)
	} else {
		c.hh -= hours
	}
	return c
}

func getTimeByMin(minutes int) (hours int, minutesLeft int) {
	hours = (minutes / 60) % 24
	minutesLeft = minutes % 60
	return
}

func getHours(hour int) int {
	return hour % 24
}
