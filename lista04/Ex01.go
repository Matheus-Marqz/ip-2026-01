package main
import f "fmt"
var b, e float64
func main() {
    f.Print("Digite dois valores, uma base e um expoente: ")
    f.Scan(&b, &e)
    r := recursao(b, e)
    f.Print("O resultado dessa exponenciação é: ")
    f.Print(r)
}
func recursao(x float64, n float64) float64{
    if n == 1 {
        return x
    }
    return x*recursao(x,n-1)
}
