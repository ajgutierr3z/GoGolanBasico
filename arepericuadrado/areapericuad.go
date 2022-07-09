package main
import (
    "fmt"
    "strconv"
)
func main () {
	var lado string
	var l, area, perimetro int
	fmt.Println("Escribe el valor del lado del cuadrado:")
	fmt.Scanln(&lado)
	l,_ = strconv.Atoi(lado)
	area = l * l
	perimetro = l * 4
	fmt.Println("El area del cuadrado es: ",area)
	fmt.Println("El perimetro del cuadro es: ",perimetro)
}
