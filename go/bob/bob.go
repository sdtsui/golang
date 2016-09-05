package bob

import (
  "strings"
)

const testVersion = 2

func Hey(s string) (response string) {
  switch s = strings.TrimSpace(s); {
  case silent(s):
    response = "Fine. Be that way!"
  case shouting(s):
    response = "Whoa, chill out!"
  case question(s):
    response = "Sure."
  default:
    response = "Whatever."
  }
  return
}

func silent(s string) bool {
  return s == ""
}

func question(s string) bool {
  return strings.HasSuffix(s, "?")
}

func shouting(s string) bool {
  return s == strings.ToUpper(s) && strings.ToUpper(s) != strings.ToLower(s)
}
// switch case, short functions
// named return