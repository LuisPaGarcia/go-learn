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
  PageTitle string
  PageTodos  []Todo
}

var todos = []Todo{}

func main(){
  http.HandleFunc("/", getTodos)
  http.HandleFunc("/todos/", getTodos)
  http.HandleFunc("/add-todo/", addTodo)
  fmt.Println("Server is running at :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))

}

func addTodo(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  todo := Todo{
    Title: r.FormValue("title"), // name att from html input
    Content: r.FormValue("content"),
  }

  todos = append(todos, todo)
  log.Print(todos)
  http.Redirect(w, r, "/", http.StatusSeeOther)

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