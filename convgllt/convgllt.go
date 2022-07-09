package main
import (
	"fmt"
	"strconv"
)
func main () {
	var galon string
	var lts float64
	fmt.Println("Escribe la cantidad de galones a convertir a litros")
	fmt.Scanln(&galon)
	lts,_ = strconv.ParseFloat(galon,64)
	lts = lts / 0.21997
	fmt.Println("La cantidad de litros es: ",lts)
}