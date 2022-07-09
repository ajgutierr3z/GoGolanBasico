package main
import (
    "fmt"
    "strconv"    
)
func main() {
    const pi = 3.1416
    var radio string
    var peri,r, area float64
    fmt.Println("Escrie el radio del circulo")
    fmt.Scanln(&radio)
    r,_= strconv.ParseFloat(radio,64)
    peri = 2*pi*r
    area = (r*r)*pi
    fmt.Println("El perimetro del circulo es: ",peri)
    fmt.Println("El area del circulo es: ",area)
}
