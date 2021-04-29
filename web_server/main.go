package main

import (
  "fmt"
  "net/http"
  "log"
  "html/template"
)

type Todo struct {
  Title, Content string
}

type PageVariables struct {
  PageTitle,PageTodos string
}

var todos = []Todo

func main(){
  http.HandleFunc("/", getTodos)
  http.HandleFunc("/todos", getTodos)
  fmt.Println("Server is running at :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request){
  fmt.Println("home!")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
  pageVariables := PageVariables{
    PageTitle: "Get Todos",
    PageTodos: todos,
  }
  t, err := template.ParseFiles("index.html")

  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    log.Print("Template error string", err)
  }


  err = t.Execute(w, pageVariables)
}