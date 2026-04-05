package main

import f "fmt" 

func main() {
    var num float64
    f.Print("Digite um número: ")
    f.Scanln(&num)
    if num == 2 {
        f.Println("Inexistente")
    } else {
         f.Printf("O valor de f(x) = 8/(2-x) é %.1f\n", 8/(2-num))
    }
    f.Println("Fim do código!")
}
