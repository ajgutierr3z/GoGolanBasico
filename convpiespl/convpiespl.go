package main
import (
	"fmt"
	"strconv"
)
func main () {
	var pies string
	var pl float64
	fmt.Println("Escribe la cantidad de pies a convertir")
	fmt.Scanln(&pies)
	pl,_ =strconv.ParseFloat(pies,64)
	pl = pl * 12.00
	fmt.Println("El total de pulgadas es: ",pl)
}
