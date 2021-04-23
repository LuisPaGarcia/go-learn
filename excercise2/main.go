package main

import "fmt"

var pets map[string]string = map[string]string{
  "darwin": "dog",
  "gary": "cat",
  "chowder": "turtle",
  "chela": "dog",
}

var groseries = []string {
  "papitas",
  "aguitas",
  "brownies",
  "agua pura",
}


func main(){
  fmt.Println(average([5]float64{1.1,2.2,3.3,4.4,5.5}))
  fmt.Println(ifPetExist("darwin"))
  fmt.Println(addingGroseries("juguito", "chocolate", "una picsa"))
}

// Calculate and return the average of 5 float numbers as array
func average(args [5]float64) float64 {
  var total = 0.0

  for _, value := range args {
    total += value
  }

  return total / float64(len(args))
}

// Validates if the pet exist in the map of pets
func ifPetExist(nameOfPet string) bool {
  _, ok := pets[nameOfPet]
  return ok
}

// Add some N newGroseries to the list of groseries
func addingGroseries(newGroseries ...string) []string {
  result := groseries
  for _, grosery := range newGroseries {
    result = append(result, grosery) 
  }
  return result
}