package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
  lengthA := len(a)
  lengthB := len(b)

  if lengthA != lengthB {
    return int(-1), errors.New("Disallow")
  }

  distance := 0

  for i := 0; i < lengthA; i++ {
    if (a[i] != b[i]) {
      distance++
    }
  }

  return distance, nil
}

// p("Len: ", len("hello"))
// p("Char:", "hello"[1])