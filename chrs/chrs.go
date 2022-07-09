package main
import (
    "fmt"
    "strconv"
)
func main () {
    var hactual, hfutura string
    var ha, hf int
    fmt.Println("Escribe el valor entero de la hora")
    fmt.Scanln(&hactual)
    ha,_ = strconv.Atoi(hactual)
    fmt.Println("Escribe la cantidad de horas")
    fmt.Scanln(&hfutura)
    hf,_ = strconv.Atoi(hfutura)
    hf = hf + ha
    fmt.Println("la hora sera: ", hf)
}
