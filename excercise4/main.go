package main

import "fmt"

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

  desc := user.describe()

  descGroup := group.describe()

  fmt.Println(desc)
  fmt.Println(descGroup)
}