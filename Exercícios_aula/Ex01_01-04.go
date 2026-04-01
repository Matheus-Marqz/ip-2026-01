package main

import f "fmt"

func sum(a, b float64) float64{
    sum := a + b
    return sum
}

func main() {
  var n1, n2 float64
  f.Print("Digite um número: ")  
  f.Scan(&n1)
  f.Print("Digite outro número: ")
  f.Scan(&n2)
  s := sum(n1, n2)
  f.Printf("A soma desses dois números é %.1f",s)
}
