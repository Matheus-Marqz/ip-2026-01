package main

import f "fmt"

var (
    dia, mes, ano int
    meses = [12]string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
)
func main() {
    f.Print("Digite um dia, um mês e um ano (01 02 2026): ")
    f.Scanln(&dia, &mes, &ano)
    f.Printf("Dia %d de %s de %d", dia, meses[mes-1], ano)
}
