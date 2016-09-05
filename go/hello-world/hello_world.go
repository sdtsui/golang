package hello

import "unicode/utf8"

const testVersion = 2

//Says Hello World
func HelloWorld(input string) string {
  if utf8.RuneCountInString(input) != 0 {
    return "Hello, " + input + "!"
  }

  return "Hello, World!"
}
