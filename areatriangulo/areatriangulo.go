package main
import (
	"fmt"
	"strconv"
)
func main () {
	var altura, base string
	var a,b, area int
	fmt.Println("Escribe el valor entero de la base del triangulo")
	fmt.Scanln(&base)
	b,_ = strconv.Atoi(base)
	fmt.Println("Escribe el valor entero de la altura")
	fmt.Scanln(&altura)
	a,_ = strconv.Atoi(altura)
	area = (b*a)/2
	fmt.Println("El valor del area es: ",area)
}
