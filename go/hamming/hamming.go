package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
  length := len(a)
  if len(a) != len(b) {
    return int(-1), errors.New("Disallow")
  }
  distance := 0
  for i := 0; i < length; i++ {
    if (a[i] != b[i]) {
      distance++
    }
  }

  return distance, nil
}

//range keyword
// for i := range a {
//       if a[i] != b[i] {
//         cnt++
//       }
//     }