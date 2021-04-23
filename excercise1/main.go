package main

import "fmt"

func average(float1 float32, float2 float32, float3 float32) float32 {
  return (float1+float2+float3)/3
}

func variadicAverage(args ...float32) float32 {
  var accum float32

  for _, value := range args {
    accum += value
  }
  return accum / float32(len(args))
}

func main() {
  // averageResult := average(1.1,2.2,3.3)
  averageVariadicResult := variadicAverage(1.0, 2.0, 3.0, 4.0)
  fmt.Println(averageVariadicResult)
}