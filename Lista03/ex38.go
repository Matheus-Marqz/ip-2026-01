package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var cpf string
	var somenteDigitos string
	var soma, resto, digito1, digito2 int

	fmt.Print("Digite o CPF: ")
	fmt.Scan(&cpf)

	for _, c := range cpf {
		if unicode.IsDigit(c) {
			somenteDigitos = somenteDigitos + string(c)
		}
	}

	if len(somenteDigitos) != 11 {
		fmt.Println("CPF invalido")
		return
	}

	if strings.Repeat(string(somenteDigitos[0]), 11) == somenteDigitos {
		fmt.Println("CPF invalido")
		return
	}

	soma = 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma = soma + int(somenteDigitos[i]-'0')*peso
		peso--
	}

	resto = soma % 11

	if resto < 2 {
		digito1 = 0
	} else {
		digito1 = 11 - resto
	}

	soma = 0
	peso = 11

	for i := 0; i < 10; i++ {
		var valor int

		if i == 9 {
			valor = digito1
		} else {
			valor = int(somenteDigitos[i] - '0')
		}

		soma = soma + valor*peso
		peso--
	}

	resto = soma % 11

	if resto < 2 {
		digito2 = 0
	} else {
		digito2 = 11 - resto
	}

	if digito1 == int(somenteDigitos[9]-'0') && digito2 == int(somenteDigitos[10]-'0') {
		fmt.Println("CPF valido")
	} else {
		fmt.Println("CPF invalido")
	}
}
