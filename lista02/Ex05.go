package main

import f "fmt" 

func main() {
    var num int
    f.Print("Digite um número inteiro: ")
    f.Scanln(&num)
    if num % 5 == 0 {
        f.Printf("O número %d é divisível por 5\n", num)
    } else {
        f.Printf("O núemero %d não é divisível por 5\n", num)
    }
    f.Print("Fim do código")
}
