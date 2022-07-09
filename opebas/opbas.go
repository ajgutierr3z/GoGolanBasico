package main
import (
	"fmt"
	"strconv"
)
func main () {
	var num1, num2 string
	var n1, n2 int
	fmt.Println("Escribe el primer valor entero")
	fmt.Scanln(&num1)
	n1,_ = strconv.Atoi(num1)
	fmt.Println("Escribe el segundo valor entero")
	fmt.Scanln(&num2)
	n2,_ = strconv.Atoi(num2)
	fmt.Println("La suma es: ",n1+n2)
	fmt.Println("la resta es: ", n1 - n2)
	fmt.Println("la multiplicación es: ",n1 * n2)
	fmt.Println("la división es: ",n1 / n2)
}
