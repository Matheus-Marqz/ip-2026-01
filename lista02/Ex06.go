package main

import f "fmt" 

func main() {
    var num1, num2 int
    f.Print("Digite um número inteiro: ")
    f.Scanln(&num1)
    f.Print("Digite outro número inteiro: ")
    f.Scanln(&num2)
    if num1 % num2 == 0 {
        f.Printf("O número %d é divisível pelo número %d\n", num1, num2)
    } else {
        f.Printf("O número %d NAO é divisível pelo número %d\n", num1, num2)
    }
    f.Println("Fim do código!")
}
