package main

import "fmt"

func main() {
    var n1, sum, aux int
    fmt.Println("Quantos números para a média? ")
    fmt.Scanln(&n1)
    for i := 0; i < n1; i++ {
        fmt.Printf("O %d°número é: ", i+1)
        fmt.Scanln(&aux)
        sum = sum + aux
    }
    fmt.Printf("A média é: %.1f e a soma é %.1f e o %.1f", float64(sum)/float64(n1), float64(sum), float64(n1))
}
