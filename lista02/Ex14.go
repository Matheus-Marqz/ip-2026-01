package main

import f "fmt"

var (
    preco float64   
    ar, pintura, vidro, direcao string
)
func main() {
    f.Print("Qual o preço inicial do carro? ")
    f.Scan(&preco)
    f.Print("Você deseja ar condicionado por mais R$ 1750,00?(s/n) ")
    f.Scan(&ar)
    f.Print("Você deseja Pintura metálica por mais R$ 800,00?(s/n) ")
    f.Scan(&pintura)
    f.Print("Você deseja Vidro elérico por mais R$ 1200,00?(s/n) ")
    f.Scan(&vidro)
    f.Print("Você deseja Direção hidráulica por mais R$ 2000,00?(s/n) ")
    f.Scan(&direcao)
    if ar == "s" || ar == "S"{
        preco += 1750
    }
    if pintura == "s" || pintura == "S"{
        preco += 800
    }
    if vidro == "s" || vidro == "S"{
        preco += 1200
    }
    if direcao == "s" || direcao == "S"{
        preco += 2000
    }
    f.Printf("O valor total do carro será de %.2f.\n", preco)
}
