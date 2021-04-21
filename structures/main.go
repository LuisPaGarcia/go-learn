package main

import "fmt"


// Retorno de valores directamente
func printAge(age int) int {
  return age*2
}

// retorno complejo de variables (mas chilero, refactorizado)
func printAgeComplex1() (edadDeAle int, edadDeMochy int) {
  edadDeAle = 23
  edadDeMochy = 25
  return 
}

// manejo de numero desconocido de parametros
func printAgesConUnknownParams (ages ...int) {
  for _, age := range ages {
    fmt.Println(age)
  }
}

func main() {
  age0 := printAge(9)
  fmt.Println(age0)

	age1, age2 := printAgeComplex1()
  fmt.Println(age1, age2)

  printAgesConUnknownParams(1,2,3,4,5,6,7,8)
  printAgesConUnknownParams(100)

}
