package clock

import "fmt"

const (
  testVersion     = 4
  MINUTES_PER_DAY = 60 * 24
)

type Clock int

func New(hour, minute int) Clock {
  mins := (hour*60 + minute) % MINUTES_PER_DAY
  for mins < 0 {
    mins += MINUTES_PER_DAY
  }
  return Clock(mins) // create new
}

func (c Clock) String() string {
  hours := c / 60
  minutes := c % 60
  return fmt.Sprintf("%02d:%02d", hours, minutes) //https://golang.org/pkg/fmt/
}

func (c Clock) Add(minutes int) Clock {
  totalMins := int(c) + minutes
  return New(totalMins/60, totalMins%60)
}
