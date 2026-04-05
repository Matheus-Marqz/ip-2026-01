package main

import f "fmt"

import "math"

var (
    A, B, C float64
)
func main() {
    f.Print("Qual o valor do coeficiente A?  ")
    f.Scan(&A)
    f.Print("Qual o valor do coeficiente b?  ")
    f.Scan(&B)
    f.Print("Qual o valor do coeficiente C?  ")
    f.Scan(&C)
    delta := B*B -4*A*C
    r1 := (-B + math.Sqrt(delta))/2*A
    r2 := (-B - math.Sqrt(delta))/2*A
    if delta == 0 {
        f.Println("Raiz Única")
        f.Printf("O valor da raíz é %.1f\n", r1)
    } else if delta > 0{
        f.Println("Raízes distintas")
        f.Printf("O valor das raízes é %.1f e %.1f\n", r1, r2)
    } else {
        f.Println("Raízes imaginárias")
        f.Printf("Não há valor real para as raízes")
    }
}
