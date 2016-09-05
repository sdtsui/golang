package gigasecond

import "time"

const (
  testVersion = 4
  GIGASECOND = 1000000000
)

func AddGigasecond(inputTime time.Time) time.Time {
  newTime := inputTime.Add(time.Duration(GIGASECOND)*time.Second)
  return newTime
}

// seconds := 10
// fmt.Print(time.Duration(seconds)*time.Second) // prints 10s