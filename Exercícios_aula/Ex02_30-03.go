package main

import f "fmt"

func main() {
	//ler as 5 notas
	var vetor [5]int
	var sum int
	for i:=0; i<len(vetor); i++ {
	    f.Printf("Coloque o valor %d: ", i+1)
	    f.Scan(&vetor[i])
	    //soma as notas
	    sum = sum + vetor[i]
	}
	//printa as notas e a soma
	f.Print("Os valores são: \n")
	for _, v:=range vetor {
	    f.Printf("%d\n", v)
	}
	f.Printf("A soma desses valores é: %d", sum)
}
