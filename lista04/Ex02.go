package main

import f "fmt"

var y, d, x float64
var a []float64

func main() {
    f.Print("Digite o ultimo valor do array: ")
    f.Scan(&y)

    f.Print("Digite a distancia entre dois desses valores: ")
    f.Scan(&d)

    x = 0

    for x <= y {
        a = append(a, x)
        x = x + d
    }

    f.Println("Array:", a)
    f.Println("Soma:", recursao(a, 0))
}

func recursao(x []float64, n int) float64 {
    if n >= len(x) {
        return 0
    }
    return x[n] + recursao(x, n+1)
}
