package main
import (
	"fmt"
	"strconv"
)
func main () {
	var pies string
	var mts float64
	fmt.Println("Escribe el numero de pies a convertir")
	fmt.Scanln(&pies)
	mts,_ = strconv.ParseFloat(pies,64)
	mts = mts /3.2808
	fmt.Println("La cantidad de metros es: ",mts)
}
