package main

import "fmt"

func main() {
    var n1, n2, p1, p2 float64
    fmt.Println("Digite dois números: ")
    fmt.Scanln(&n1)
    fmt.Scanln(&n2)
    fmt.Println("Digite os dois pesos: ")
    fmt.Scanln(&p1)
    fmt.Scanln(&p2)
    fmt.Printf("A média ponderada entre eles é: %.1f", (n1*p1 + n2*p2)/(p1+p2))
}
