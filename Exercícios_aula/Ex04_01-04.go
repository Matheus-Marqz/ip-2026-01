package main

import f "fmt"

func fatorial(a int) int{
    multiplicacao := a
    for a > 1 {
        a = a-1
        multiplicacao = multiplicacao*a
    }
    return multiplicacao
}
func main() {
  var n1 int
  f.Print("Digite um número: ")  
  f.Scan(&n1)
  fatorial := fatorial(n1)
  f.Printf("O fatorial de %d é: %d", n1, fatorial)
}
