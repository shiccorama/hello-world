package main
import "fmt"

type laptop struct{
processor int
color string
ram float64
}

func main(){

var mine=laptop{

processor:3,
color:"grey",
ram:8.2,
}

fmt.Println(mine.ram)
fmt.Println(mine)


}