package main

import f "fmt" 
import "math"

func main() {
    var num float64
    f.Print("Digite um número real: ")
    f.Scanln(&num)
    if num >= 0 {
        f.Printf("A raiz quadrada desse número é %.2f\n", math.Sqrt(num))
    } else {
        f.Printf("O quadrado desse número negativo é: %.1f\n", num*num)
    }
    f.Println("Fim do código!")
    
}
