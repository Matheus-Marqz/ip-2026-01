
package main

import f "fmt"

func media(a, b, c float64) float64{
    m := (a+b+c)/3
    return m
}
func main() {
  var n1, n2, n3 float64
  f.Print("Digite um número: ")  
  f.Scan(&n1)
  f.Print("Digite outro número: ")
  f.Scan(&n2)
  f.Print("Digite mais um número: ")
  f.Scan(&n3)
  m := media(n1, n2, n3)
  f.Printf("A média desses valores é: %.1f",m)
}
