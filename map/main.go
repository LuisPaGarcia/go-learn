// Un MAP es un equivalente a un objeto en JS

package main

import "fmt"

func main() {
  comoFuncionaUnMap()
}

func comoFuncionaUnMap(){
  // En maps (comparandolo con slice, no debemos establecer el tamaño initial)
  var userEmails map[int]string = make(map[int]string)

  userEmails[1] = "test@domain.com"
  userEmails[2] = "tost@damain.com"

  fmt.Println(userEmails)

  // Otra forma de inicializar un map
  userEmails2 := map[int]string{
    1: "tist@diminn.com",
    2: "tost@domoin.com",
  }
  // Y podemos editar un item del map sin pedos
  userEmails2[1] = "tust@dumuin.cum"
  
  fmt.Println(userEmails2)

  // POdemos extraer la data de un map asignandola a una variable
  user1, isOK := userEmails2[1]
  fmt.Println(user1)

  // va, asignar un item de map retorna 2 valores 🤯
  user5, ok := userEmails2[5]
  // `ok` en este caso va a indicar si el elemento existe
  fmt.Println(user5, ok) // false si no existe
  fmt.Println(user1, isOK) // true si si

  // Usaremos el `ok` dentro de un if, 
  if _, ok := userEmails[1]; ok { // Si OK es true, haga lo que esta dentro
    fmt.Println("Email exists!")
  } else {
    fmt.Println("Email doesn't exist")
  }

  // Y para eliminar un nodo de un map, usamo `delete`
  delete(userEmails, 1) // Mandamos el map y el key name para eliminar

  fmt.Println(userEmails)


}