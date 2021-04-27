package main

import "fmt"

type User struct {
  fname, lname, email string
}

  var user = User{
    fname: "luispa",
    lname: "garcia",
    email: "hola@luispa.im",
  }


func main() {

  updateEmail(&user)
  fmt.Println(user)
}

func updateEmail(u *User) {
  u.email = "paulypeligroso1@gmail.com"
}