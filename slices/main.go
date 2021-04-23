package main

import "fmt"

func main(){
//  comoUsarSlices()
subSelectUnArray()
}

func comoUsarSlices(){
   var array [5]int
  // Un make recibe el tipo, el tamano minimo y el tamano maximo (opdional)
  // Todos los slices son arrays behind the scenes
  var slice []int = make([]int, 5, 10)
  var slice2 []int = make([]int, 5, 12)

  array[0] = 1
  slice[0] = 1
  slice2[4] = 1

  fmt.Println(array)
  fmt.Println(slice)
  fmt.Println(slice2)
  fmt.Println(len(slice)) // conocer el largo
  fmt.Println(cap(slice)) // conocer la capacidad!
}

func subSelectUnArray() {
  frutasArray := [...]string{"pera", "manzana","tomate","aguacate"}

  splicedFruits := frutasArray[1:3] // [manzana, tomate]
  fmt.Println(splicedFruits)
  fmt.Println(cap(splicedFruits))
  fmt.Println(len(splicedFruits))
  // Es curioso, pero go va a aumentar la CAP 
  // de un slice basado en el tamaño original
  // y no importa que le estemos agregando items
  // Go va a allocar memoria para guardar lo que necesitemos :D
  addedFruits := append(splicedFruits, "pepino","naranja", "moritas")
  fmt.Println(addedFruits)
  fmt.Println(cap(addedFruits))
  fmt.Println(len(addedFruits))

}