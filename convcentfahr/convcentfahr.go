package main 
import (
	"fmt"
	"strconv"
)
func main () {
	var centi string
	var fahr float64
	fmt.Println("Escribe la cantidad de gradados centigrados a convertir")
	fmt.Scanln(&centi)
	fahr,_=strconv.ParseFloat(centi,64)
	fahr = (fahr * 1.8000)+32.00
	fmt.Println("la conversión a grados Fahrenheit es de:",fahr)
}
