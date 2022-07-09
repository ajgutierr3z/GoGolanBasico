package main
import (
	"fmt"
	"strconv"
)
func main () {
	var pulgada string
	var pl float64
	fmt.Println("Escribel la cantidad de pulgadas a convertir")
	fmt.Scanln(&pulgada)
	pl,_ = strconv.ParseFloat(pulgada,64)
	pl = pl / 39.370
	fmt.Println("La cantidad en metros es: ",pl)
}
