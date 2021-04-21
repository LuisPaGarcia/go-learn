package main

import (
  "fmt"
  "reflect"
  )

var name string = "Luizpa"
var unicode string = "123日本語"
var weirdChars string = "⚽😎"

func main() {
	// howToUsePrintF()
  // comoUsarReflect()
  // usandoVariablesGlobales(name)
  // ReverseString(unicode)
  // IfTrucazo()
  // usoDeSwitch()
  // usoDeForLoops()
  pruebitaDeLoops()
}

func pruebitaDeLoops(){
  var oracion string = "Esto es un ejemplo de una oracion"

  for _, letter := range oracion {
      fmt.Println(string(letter))
  }
}

func usoDeForLoops() {
  // loop parecido a js
  for i := 1; i <= 100; i++ {
    fmt.Print(i)
  }

  // Un for loop que se parece a un while
  var index int = 1
  for index <= 100 {
    fmt.Print(index)
    index += 1
  }

  var oracion string = "Esta es una oracion"
  // iterando sobre una oracion
  for index, letter := range oracion {
    fmt.Println("index:", index, "letra:", string(letter))
  }

}

func usoDeSwitch() {
  // Implementar un switch comun
  var ciudad string = "sofy"
  switch ciudad {
    case "luispa":
      fmt.Println("is luispa")
    case "andre":
      fmt.Println("is andre")
    case "ale", "sofy":
      fmt.Println("is ale or sofy")
  }

  // Evaluar un valor externo al swich
  var edad int = 10
  switch {
    case edad > 0:
      fmt.Println("mayor a 0 ")
  }

  // Como usar un default en un if
  var houseNumber int = 10
  switch {
    case houseNumber > 10:
      fmt.Println("Es mayor que uno")
    case houseNumber < 10:
      fmt.Println("Es menor que diez")
    default:
      fmt.Println("es diez")
  }

  // Cuando un switch hace match, go no sigue
  // A MENOS de que usemos fallthrough
  var number int = 9
  switch {
    case number != 10:
      fmt.Println("No es igual que diez")
      fallthrough
    case number < 10:
      fmt.Println("es menor que diez")
    case number > 10:
      fmt.Println("es mayor que diez")
    default:
      fmt.Println("Es 10")
  }
  
}

func IfTrucazo(){
  // Puedes definir en el mismo if una variable con solo este scope
  if err := ReverseString(unicode); err != unicode {
    
  }

}

func ReverseString(str string) string {
  output :=""
  for _, char := range str {
    output = string(char) + output
  }
  fmt.Println(output)
  return str
}

func usandoVariablesGlobales(aString string) string {
  fmt.Println(aString)
  return aString
}

func howToUsePrintF(){
	fmt.Printf("Hola! Soy %s y vivo en la antigua desde hace %d años, el clima es hermoso y eso es %t\n\n\n", "Luispa", 29, true)

}

func comoUsarReflect() {
  var largeInt uint64 = 10000000000000000000  
  fmt.Println("Tipo de 1: ", reflect.TypeOf(1))
  fmt.Println("Tipo de hey: ", reflect.TypeOf("hey"))
  fmt.Println("Tipo de true: ", reflect.TypeOf(true))
  fmt.Println("Tipo de 10000000000000000000: ", reflect.TypeOf(largeInt))
}

