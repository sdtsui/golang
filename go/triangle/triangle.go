package triangle

import "math"

const (
  testVersion = 2
  NaT Kind = iota //use of iota
  Equ
  Iso
  Sca
)
// can also do var NaT Kind = "not a triangle"

func KindFromSides(a, b, c float64) Kind {
  sides := [6]float64{a, b, c, a, b, c}
  if (!isTriangle(sides)) {
    return Kind(NaT)
  }
  if (a == b && b == c && a == c) {
    return Kind(Equ)
  }
  if (a == b || b == c || c == a) {
    return Kind(Iso)
  }

  return Kind(Sca)
}

func isTriangle(sides [6]float64) bool { //short helper
  for i := range sides[:3] {
    if math.IsInf(sides[i], 0) || math.IsNaN(sides[i]) || sides[i] == 0 {
      return false
    }
    sumOfLengths := sides[i] + sides[i+1]
    if (sumOfLengths < sides[i+2]) {
      return false
    }
  }

  return true
}

type Kind int

/**
 * 
 * type Kind struct {
  str string
}

var NaT Kind = Kind{str: "NaT"} // not a triangle
var Equ Kind = Kind{str: "Equ"} // equilateral
var Iso Kind = Kind{str: "Iso"} // isosceles
var Sca Kind = Kind{str: "Sca"} // scalene

 */
