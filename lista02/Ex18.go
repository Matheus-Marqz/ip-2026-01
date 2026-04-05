package main

import f "fmt"


func main() {
    var (
        dia int   
        valor, pagamento float64
        categoria string
    )
    
    f.Print("Qual o dia da semana que estamos?\n")
    f.Print("\n 1- Domingo \n 2- Segunda \n 3-Terça \n 4-Quarta \n 5-Quinta \n 6-Sexta \n 7-Sabado\n")
    f.Scan(&dia)
    f.Print("Qual o valor do DVD?(R$) ")
    f.Scan(&valor)
    if dia == 2 || dia == 3 || dia == 5 {
        pagamento = valor * 0.6
    } else {
        pagamento = valor
    }
    f.Print("Qual a categoria do DVD?(Comum ou Lançamento) ")
    f.Scan(&categoria)
    if categoria == "Lançamento" || categoria == "lançamento" {
        pagamento += valor * 0.15
    }
    f.Printf("O valor pago pelo aluguel do DVD foi de %.2f\n", pagamento)
}
