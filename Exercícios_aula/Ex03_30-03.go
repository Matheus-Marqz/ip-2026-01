package main

import f "fmt"

func main() {
	//ler as 10 valores
	var vetor [10]float64
	i:=len(vetor)
	n := 0
	for i>0 {
	    i = i - 1
	    n = n + 1
	    f.Printf("Digite o %d° valor: ", n)
	    //colocar o número na ultima posição livre do array
	    f.Scan(&vetor[i])
	}
	//Printar vetor invertido
	f.Print(vetor)
}
