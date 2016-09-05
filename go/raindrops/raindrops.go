package raindrops

import "strconv"

const testVersion = 2

func Convert(number int) string {
  reply := ""
  if (number % 3) == 0 {
    reply += "Pling"
  }
  if (number % 5) == 0 {
    reply += "Plang"
  }
  if (number % 7) == 0 {
    reply += "Plong"
  }
  if len(reply) == 0 {
    reply = strconv.Itoa(number)
  }

  return reply
}

// var words = map[int]string{
//   3: "Pling",
//   5: "Plang",
//   7: "Plong",
// }
