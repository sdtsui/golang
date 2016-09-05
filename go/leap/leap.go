package leap

import "fmt"

const testVersion = 2

// Tells us if a year is a leap year
func IsLeapYear(year int) bool {
  evenlyDivisibleByFour := (year % 4) == 0
  evenlyDivisibleByOneHundred := (year % 100) == 0
  evenlyDivisibleByFourHundred := (year % 400) == 0

  if evenlyDivisibleByFour {
    if evenlyDivisibleByOneHundred {
      if evenlyDivisibleByFourHundred {
        return true
      }
      return false
    }
    return true
  }

  return false
}
