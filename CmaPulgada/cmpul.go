package main
import (
   "fmt"
   "strconv"
)
func main() {
   var cm string
   var cmapl float64
   const cmt = 2.56
   fmt.Println("Escribe los centimetros a convertir:")
   fmt.Scanln(&cm)
   cmapl,_= strconv.ParseFloat(cm,64)
   cmapl = cmapl / cmt
   fmt.Println("Las pulgadas son: ",cmapl)
}
