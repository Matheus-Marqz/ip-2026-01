package main

import f "fmt" 

func main() {
    var num int
    f.Print("Digite uma idade: ")
    f.Scanln(&num)
    switch {
        case num >= 0 && num <= 2:
            f.Println("Recém-nascido")
        case num >= 3 && num <= 11:
            f.Println("Criança")
        case num >= 12 && num <= 19:
            f.Println("Adolescente")
        case num >= 20 && num <= 55:
            f.Println("Adulto")
        case num > 55:
            f.Println("Idoso")
        default:
            f.Println("Idade inválida!")
    }
    f.Println("Fim do código!")
}
