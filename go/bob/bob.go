package bob

import (
  "strings"
  "fmt"
)

const (
  testVersion = 2
  silenceReply = "Fine. Be that way!"
  questionReply = "Sure."
  shoutReply = "Whoa, chill out!"
  defaultReply = "Whatever."
)


func Hey(s string) (response string) {
  s = strings.TrimSpace(s)

  if (s == "") {
    return silenceReply
  }

  if (s == strings.ToUpper(s) && s != strings.ToLower(s)) {
    return shoutReply
  }

  if (strings.HasSuffix(s, "?")) {
    return questionReply
  }

  return defaultReply
}

// alternate: switch case, short functions
// named return
// 
// question--ends  with a ?
// silence--empty
// caps--