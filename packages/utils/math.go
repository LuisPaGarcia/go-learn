package utils

import "fmt"

func printNum(number int) {
  fmt.Println("Current number:", number)
}

// Add sums multiple integers
func Add(nums ...int) int{
  result := 0
  for _, num := range nums{
    printNum(num)
    result += num
  }

  return result
}