package main
import "fmt"

type car struct{
color string
Speed float64
Price int
}

func main(){

var vw=car {
color:"blue",
Speed:88.88,
Price:1500,}

GetPrice(vw)

         fmt.Println("out of Discount")}

func GetPrice(vw car){
    vw.Price=14000
	fmt.Println(vw.Price)}
