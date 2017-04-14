package main

import "fmt"

func main() {
  var scores [5]float64
  scores[0] = 98
  scores[1] = 91
  scores[2] = 77
  scores[3] = 82
  scores[4] = 83

  var total float64 = 0
  for _, value := range scores {
    total += value
  }
  fmt.Println(total / float64(len(scores)))
}
