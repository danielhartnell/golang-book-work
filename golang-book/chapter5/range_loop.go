package main

import "fmt"

func main() {
  var x [3]int
  x[0] = 2
  x[1] = 4
  x[2] = 6

  for _, value := range x {
    fmt.Println(value)
  }
}
