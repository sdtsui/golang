How to declare a map:
```
var words = map[int]string{
  3: "Pling",
  5: "Plang",
  7: "Plong",
}
```

using the range keyword:
```
// for i := range a {
//       if a[i] != b[i] {
//         cnt++
//       }
//     }
```


String Manipulation Notes, from Bob:

  HasSuffix(s, "?")  >  Last Char: `s[len(s)-1]`
  string methods:
    -ContainsAny
    -ToUpper
    -TrimSpace

  Top-level constants, to be used again


  unicode package:
    unicode.isUpper
    unicode.isLower

  declare maps of different types:
  map[string]string{
    "question":       "Sure.",
    "yell":           "Whoa, chill out!",
    "nothing-to-say": "Fine. Be that way!",
    "anything-else":  "Whatever.",
  }


  func any(items string, test func(rune) bool) bool {
    for _, item := range items {
      if test(item) {
        return true
      }
    }
    return false
  }

  func IsUpper
  func IsUpper(r rune) bool
  IsUpper reports whether the rune is an upper case letter.

