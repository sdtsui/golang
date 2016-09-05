package triangle

const testVersion = 2

// Code this function.
func KindFromSides(a, b, c float64) Kind

// Notice it returns this type.  Pick something suitable.
type Kind

// Pick values for the following identifiers used by the test program.
NaT // not a triangle
Equ // equilateral
Iso // isosceles
Sca // scalene

// Organize your code for readability.

// find number of sides, and side lengths
// 
// edge case: lines can't match?
//  {NaT, 0, 0, 0},    // zero length
  // {NaT, 3, 4, -5},   // negative length
  // {NaT, 1, 1, 3},    // fails triangle inequality