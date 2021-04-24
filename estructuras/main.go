package main

import "fmt"

// O podemos agrupar los mismos tipos en una misma linea

//   type User struct {
//     ID int
//     Email, FirstName, LastName string
//   }

// User representa un usuario
type User struct {
  ID int
  Email string
  FirstName string
  LastName string
}

type Group struct {
  role string
  users []User
  newestUser User
  spaceAvailable bool
}


func main() {
  // Se usa un struct como un tipo, y se llena como un slice o un array
  user := User{
    ID: 1,
    Email: "hola@luispa.im",
    FirstName: "Luispa",
    LastName: "Garcia",
  }

  user2 := User{
    ID: 2,
    Email: "andrius@gmail.com",
    FirstName: "Andrea",
    LastName: "Martinez",
  }

  // tambien se pueden acceder a las propiedades como props de js
  fmt.Println(user.ID)
  fmt.Println(user.Email)
  fmt.Println(user.FirstName)
  fmt.Println(user.LastName)
  fmt.Println(user)

  group := Group{
    role: "admin",
    users: []User{user},
    newestUser: user2,
    spaceAvailable: true,
  }

  fmt.Println(describeUser(user))
  fmt.Println(describeGroup(&group))
  fmt.Println(group.spaceAvailable)
}

func describeUser(u User) string {
  description := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
  return description  
}

func describeGroup(g *Group) string {
  if len(g.users) <= 2 {
    g.spaceAvailable = false
  }
  description := fmt.Sprintf("The user group has %d users. The newest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
  return description
}