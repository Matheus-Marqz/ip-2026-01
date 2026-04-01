
package main

import f "fmt"

func maior(a, b, c float64) [3]float64{
    var aux float64
    if a > b {
        aux = a
        a = b
        b = aux
    } 
    if b > c {
        aux = b
        b = c
        c = aux
    }
    if a > b {
        aux = a
        a = b
        b = aux
    }
    return [3]float64{a,b,c}
}
func main() {
  var n1, n2, n3 float64
  f.Print("Digite um número: ")  
  f.Scan(&n1)
  f.Print("Digite outro número: ")
  f.Scan(&n2)
  f.Print("Digite mais um número: ")
  f.Scan(&n3)
  m := maior(n1, n2, n3)
  f.Printf("A ordem crescente é: %.1f, %.1f, %.1f",m[0], m[1], m[2])
}
