package main

import f "fmt"

func main() {
    var num int 
    f.Println("Digite um número inteiro: ")
    f.Scanln(&num)
    if num % 2 == 0 {
        f.Println("Este número é par")
    } else {
        f.Println("Esse número é ímpar")
    }
    f.Println("Fim do código.")
}
