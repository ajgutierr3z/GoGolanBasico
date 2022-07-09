package main
import (
    "fmt"
    "strconv"
    "math"
)
func main (){
    var ca, cb, hipo float64
    var Ca, Cb string
    fmt.Println("Escribe el valor del cateto a:")
    fmt.Scanln(&Ca)
    ca,_=strconv.ParseFloat(Ca,64)
    fmt.Println("Escribe el valor del cateto b:")
    fmt.Scanln(&Cb)
    cb,_=strconv.ParseFloat(Cb,64)
    ca= ca * ca
    cb= cb * cb
    hipo = math.Sqrt(ca+cb)
    fmt.Println("El calculo de la hipotenusa es: ",hipo)
}
