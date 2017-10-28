package main

import "fmt"

func main() {
  var n int
  fmt.Scanf("%v\n", &n)
  e := readInput(n)
  result := sumArray(e)
  fmt.Printf("%v", result)
}

func readInput(n int) []int{
  r := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%v", &r[i])
  }
  return r
}

func sumArray(a []int) int {
  r := 0
  for _,i := range a {
    r = r + i
  }
  return r
}