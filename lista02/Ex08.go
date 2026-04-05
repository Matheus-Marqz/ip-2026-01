package main

import f "fmt" 

func main() {
    var num int
    f.Print("Digite um número inteiro: ")
    f.Scanln(&num)
    if num > 20 && num < 90 {
        f.Printf("O número %d está entre 20 e 90\n", num)
    } else {
        f.Printf("O núemro %d não está entre 20 e 90\n", num)
    }
    f.Println("Fim do código!")
}
