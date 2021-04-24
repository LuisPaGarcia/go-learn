package main

import "fmt"

func main(){
  // pointersExample1()
  pointersExample2()

}

func pointersExample1(){
  var name string
  var namePointer *string // Significa POINTER TYPE

  fmt.Println("Name:", name)
  fmt.Println("Name *:", namePointer)
}

func pointersExample2() {
  var name string
  var namePointer *string

  name = "Travis"
  namePointer = &name // Obtiene la direccion de memoria 
  var nameValue = *namePointer // Obtiene el valor de la direccion de memoria

  fmt.Println("Name:", name)
  fmt.Println("Name *:", namePointer)
  fmt.Println("NameValue *:", nameValue)
}