package main

import ("fmt"
"strings")

func main(){
  // pointersExample1()
  // pointersExample2()
  pointersExample3()

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

// Cuando tenemos que tratar con un chunk de data grande
// es mejor no pasar todo el chunk
// solo con pasar la direccion de memoria es 
// mas eficiente
func pointersExample3() {
  name := "jorge"
  toUpper(&name) 
  fmt.Println(name)
}

func toUpper(str *string){
  *str = strings.ToUpper(*str)
}
