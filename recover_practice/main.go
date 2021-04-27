package main

import (
	"fmt"
)

// Ejemplo de defer (Aun hay que practicar)
func hasCosas() {
	defer fmt.Println("First line!")
	defer fmt.Println("Second line!")
	fmt.Println("Third line!")
}

func recoverFromPanic() {
  if r := recover(); r != nil {
	  fmt.Println("we paniued! but everything's alright")
	  fmt.Println("Third line!")

  }
}

func panicAtTheGosco() {
  defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("PANIC")
		}
	}
}

func main() {
	panicAtTheGosco()
}
