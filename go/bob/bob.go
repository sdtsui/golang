package bob // package name must match the package name in bob_test.go

import "fmt"
import "strings"

const testVersion = 2 // same as targetTestVersion

func Hey(message string) string {
  capChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  // lowChars := "abcdefghijklmnopqrstuvwxyz"
  // numbers := "1234567890"
  lowUmaluts := "ä"
  capUmaluts := "Ä"
  // trims := numbers + " "

  // message = strings.Trim(message, trims)
  // fmt.Println("trimmedmessage:", message)


  numCaps := 0
  for i := range message {
    fmt.Println("message i :", message[i])
    if (strings.ContainsAny(message[i:i+1], capChars)) {
      numCaps++
      fmt.Println("caps")
    }

  }
  percentageCaps := float64(numCaps)/float64(len(message)) || 0
  fmt.Println("numCaps, Len", numCaps, len(message))
  fmt.Println("percentageCaps:", percentageCaps)
  if (float64(percentageCaps) > float64(0.5)) {
    return "Whoa, chill out!"
  }

  if (strings.ContainsAny(message, capUmaluts)) {
    return "Whoa, chill out!"
  }

  if (strings.ContainsAny(message, lowUmaluts)) {
    return "Whatever."
  }


  length := int(len(message))
  fmt.Println("length:", length, message[length-1:], message)
  if (message[length-1:] == "?") {
    return "Sure."
  }
  return "Whatever."
}
