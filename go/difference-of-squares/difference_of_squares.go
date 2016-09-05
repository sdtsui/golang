package diffsquares

import "math"

func SquareOfSums(n int) int {
  sums := float64(0)
  for i := 0; i<= n; i++ {
    sums += float64(i)
  }
  return int(math.Pow(sums, float64(2)))
}

func SumOfSquares(n int) int {
  sumOfSquares := float64(0)
  for i := 0; i <= n; i++ {
    sumOfSquares += math.Pow(float64(i), float64(2))
  }
  return int(sumOfSquares)
}

func Difference(n int) int {
  return int(SquareOfSums(n) - SumOfSquares(n))
}
