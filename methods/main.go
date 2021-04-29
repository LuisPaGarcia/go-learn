package main

import "fmt"

type User struct {
  ID int
  Email string
  FirstName string
  LastName string
}

// Funcion normal para pasarale como argumento un User
func describeUser(u User) string {
  description := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
  return description  
}

// Pero, para un metodo es como atacharlo a la estructura User, como prototype

func (u *User) describe() string { // se define como un prototype en la estructura
  description := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
  return description     
}

func main(){
  user := User {
    ID: 1, 
    FirstName: "LuisPa",
    LastName: "Garcia",
    Email: "lui@gar.com",
  }
  // Asi se llama la funcion
  // desc := describeUser(user)

  // Asi se llama el método
  desc := user.describe()
  // Y ambos hacen lo mismo :O

  fmt.Println(desc)
}