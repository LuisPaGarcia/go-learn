package main

import(
  "fmt"
  "errors"
  "os"
)

func main(){
  num := 9
  
  if err := esMayorQueDiez(num); err != nil {
    fmt.Println(fmt.Errorf("%d no es mayor que 10", num))
  }

  if err := openFile(); err != nil {
    fmt.Println(fmt.Errorf("%v", err))
  }
}

func esMayorQueDiez(num int) error {// podemos retornar un error o nil
  if num < 10 {
    return errors.New("something bad happened")
  }
  return nil
}

func openFile() error{
  f, err := os.Open("missingFileName.txt")
  if err != nil {
    return err
  }
  defer f.Close()
  return nil
}