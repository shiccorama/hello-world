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
Price:1500,
}

GetPrice(vw)
Getcolor(vw)


fmt.Println("out of Discount")
fmt.Println(vw.Speed)

}

func GetPrice(vw car){
	fmt.Println("inside Discount")
	fmt.Println("instead of",vw.Price)
	vw.Price=14000
	fmt.Println(vw.Price)
}

func Getcolor(vw car){
	fmt.Println(vw.color)
	
}