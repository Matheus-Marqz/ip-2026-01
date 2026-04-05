package main

import f "fmt" 

func main() {
    var A, B, C, maior, menor, inter, aux int
    f.Print("Digite um número inteiro: ")
    f.Scanln(&A)
    f.Print("Digite outro número inteiro: ")
    f.Scanln(&B)
    f.Print("Digite mais um número inteiro: ")
    f.Scanln(&C)
    if A > B {
        aux = A
        A = B
        B = aux
    } 
    if B > C {
        aux = B
        B = C
        C = aux
    } 
    if A > B {
        aux = A
        A = B
        B = aux
    }
    menor = A
    inter = B
    maior = C
    f.Printf("A ordem crecente é %d, %d, %d\n", menor, inter, maior)
    f.Println("Fim do código!")
}
