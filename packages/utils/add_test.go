package utils

import "testing"


func TestAdd(t *testing.T) {
  expected := 10
  actual := Add(1,2,3,4)
  if actual != expected {
    t.Errorf("no iguales se espera %d pero se obtiene %d", expected, actual)
  }
}