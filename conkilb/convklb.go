package main
import (
	"fmt"
	"strconv"
)
func main () {
	var kilos string
	var lib float64
	fmt.Println("Introduce los kilogramaos a convertir a libras")
	fmt.Scanln(&kilos)
	lib,_ =strconv.ParseFloat(kilos,64)
	lib = lib * 2.2
	fmt.Println("El total de libras es: ",lib)

}
