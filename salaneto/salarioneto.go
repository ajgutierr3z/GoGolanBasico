package main
import (
	"fmt"
	"strconv"
)
func main () {
	var salario string
	var Salario float64
	fmt.Println("Escribe la cantidad de horas trabajadas")
	fmt.Scanln(&salario)
	Salario,_ = strconv.ParseFloat(salario,64)
	Salario = Salario * 50
	Salario = Salario + (Salario * 0.02)
	Salario = Salario - ((Salario * 0.015) + (Salario * 0.012))
	fmt.Println("El salario neto del trabajador es: ",Salario)
}
