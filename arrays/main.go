package main

import "fmt"

func main() {
	// Como crear un array vacio
	var scores [5]float64
	// Maneras de llenar un array

  // 1. Definir el tipo y con un cast definir los valores
	var notas [5]float64 = [5]float64{9, 10.1, 14.11, 11.5}
  // 2. Definir solo con un cast e inferir segun los valores
	notas2 := [5]float64{1.0, 2.0, 3.0}
  // 3. Utilizar variaditic notation para que el largo no sea definitivo
	notas3 := [...]float64{1.0, 2.0, 3.0}

	fmt.Println(scores, notas, notas2, notas3)
}
