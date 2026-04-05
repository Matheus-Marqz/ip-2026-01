package main

import f "fmt"

func main() {
	var num int
	f.Print("Digite um número: ")
	f.Scan(&num)
	if num <= 999 && num>=100 {
	    c:= (num/10)*10
	    u := (num - c)
	    d :=  num - (num/100)*100 - u
	    f.Printf("O algarismo das dezenas é %d\n", d/10)
	} else {
	    f.Print("Não tem 3 dígitos\n")
	}
	f.Print("Fim do código")
}
