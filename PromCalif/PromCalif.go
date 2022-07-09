package main
import (
    "fmt"
    "strconv"
)
func main (){
    var cal1, cal2, cal3, cal4 string
    var prom, c1, c2, c3, c4 float64
    fmt.Println("Escribe la primera calificacion")
    fmt.Scanln(&cal1)
    c1,_=strconv.ParseFloat(cal1,64)
    fmt.Println("Escribe la segunda calificacion")
    fmt.Scanln(&cal2)
    c2,_=strconv.ParseFloat(cal2,64)
    fmt.Println("Escribe la tercera calificacion")
    fmt.Scanln(&cal3)
    c3,_=strconv.ParseFloat(cal3,64)
    fmt.Println("Escribe la califiacion cuatro")
    fmt.Scanln(&cal4)
    c4,_=strconv.ParseFloat(cal4,64)
    prom=(c1+c2+c3+c4)/4
    fmt.Println("El promedio de las calificaciones es: ",prom)
}
