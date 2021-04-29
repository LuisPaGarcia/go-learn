package main

import "fmt"

type User interface{}
type Admin interface{}
type Parent interface{}

var people map[string]interface{}
// var people = map[string]interface {
//   "user": User
//   "admin": Admin
//   "parent": Parent
// }

func main(){
  
  fmt.Println(people)
}