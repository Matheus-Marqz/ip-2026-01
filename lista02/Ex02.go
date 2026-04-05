package main

import f "fmt"

func main() {
    var num int
   f.Print("Digite um número inteiro: ")
   f.Scanln(&num)
   switch {
        case num > 0:
            f.Print("Esse número é positivo\n")
        case num == 0:
            f.Print("Esse número é nulo\n")
        default:
            f.Print("Esse número é negativo\n")
   }
   f.Println("Fim do código.")
    
}
