package main
import (
        "fmt"
        "strings"
)
var nombre string
func main(){
   fmt.Println("Escribe tu nombre: ")
   fmt.Scanln(&nombre)
   nombre=strings.TrimSpace(nombre)
   fmt.Println("Hola ",nombre)
}
