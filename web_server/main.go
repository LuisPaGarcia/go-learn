package main

import (
  "fmt"
  "net/http"
  "log"
)

func main(){
  http.HandleFunc("/", home)
  http.HandleFunc("/todos", todos)
  fmt.Println("Server is running at :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request){
  fmt.Println("home!")
}

func todos(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hey!")
}