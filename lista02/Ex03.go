package main

import f "fmt"

func main() {
    var num1, num2, sum int
    f.Print("Digite um número inteiro: ")
    f.Scanln(&num1)
    f.Print("Digite outro: ")
    f.Scanln(&num2)
    sum = num1 + num2
    if sum > 20 {
        f.Printf("O valor da soma(%d) desses números é maior que 20, e acrescido 8 vale %d\n", sum, sum + 8)
    } else {
        f.Printf("O valor da soma(%d) dos números é menor ou igual a 20, e decrescido de 5 vale %d\n", sum, sum - 5)
    }
    f.Print("Fim do código!")
}
