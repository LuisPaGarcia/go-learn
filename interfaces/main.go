package main

import "fmt"
// Interfaces: Como establecer que limitantes/permisos puede tenes un struct con sus métodos atachados, y luego poder usar el interface como un type dentro de las funciones.

// Add a describer Interface
type Describer interface {
  describe() string
   
}


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

// Pero, para un metodo es como atacharlo a la estructura User, como prototype

func (u *User) describe() string { // se define como un prototype en la estructura
  description := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
  return description     
}

func (g *Group) describe() string {
  if len(g.users) <= 2 {
    g.spaceAvailable = false
  }
  description := fmt.Sprintf("The user group has %d users. The newest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
  return description
}

// DescriptorMaster usa una interfaz
func DescriptorMaster(d Describer) string {
  return d.describe()
}


func main(){
  user := User {
    ID: 1, 
    FirstName: "LuisPa",
    LastName: "Garcia",
    Email: "lui@gar.com",
  }
  
  user2 := User{
    ID: 2,
    Email: "andrius@gmail.com",
    FirstName: "Andrea",
    LastName: "Martinez",
  }

  group := Group{
    role: "admin",
    users: []User{user, user2},
    newestUser: user2,
    spaceAvailable: true,
  }
  // con metodo
  //desc := user.describe()
  //descGroup := group.describe()

  // con interfaz
  userDescriptionWithInterface := DescriptorMaster(&user)
  groupDescriptionWithInterface := DescriptorMaster(&group)
  
  /*
  fmt.Println(desc)
  fmt.Println(descGroup)
  */
  fmt.Println(userDescriptionWithInterface)
  fmt.Println(groupDescriptionWithInterface)
}