package main
import (
   "fmt"
   "strconv"
)
func main () {
   var ed string
   var edad  int
   fmt.Println("Escribe tu edad: ")
   fmt.Scanln(&ed)
   edad,_ = strconv.Atoi(ed)
   fmt.Println("Decadas vividas: ", edad/10)
   fmt.Println("Lustros vividos: ", edad/5)
   fmt.Println("Años vididos: ",edad)
   fmt.Println("Meses vividos: ",edad *12)
   fmt.Println("Semana vividas: ",edad * 52)
   fmt.Println("Dias vividos: ",edad * 365)
   fmt.Println("Horas vividas: ",(edad *(365*24)))
   fmt.Println("Minutos vividos: ",(edad * (365*24)*60))
}
