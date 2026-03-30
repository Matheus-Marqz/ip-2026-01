package main

import f "fmt"

func main() {
	//ler as 3 notas
	var vetor [3]float64
	var sum float64
	for i := 0; i < 3; i++ {
	    f.Printf("Coloque a nota %d: ", i+1)
	    f.Scan(&vetor[i])
	    //somar
	    sum += vetor[i]
	}
	//printar as notas e a soma
	for i, v:=range vetor {
	    f.Printf("A nota %d é %.1f\n", i+1, v)
	}
	f.Printf("A soma é %.1f", sum)
}
