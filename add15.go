package main
import "fmt"

type car struct{
NumWheel int
Color string
Speed int
}

func main(){


var mycar=car{
NumWheel:4,
Color:"blue",
Speed:10,
}

GetColor (mycar)
fmt.Println("outside getcolor func")
fmt.Println(mycar.Color)

GetSpeed (mycar)
fmt.Println("outside getspeed func")
fmt.Println(mycar.Speed)

}

func GetColor(mycar car){
   fmt.Println("inside getcolor func")
   fmt.Println(mycar.Color)
   mycar.Color="Red"
   fmt.Println(mycar.Color)
   	
}

func GetSpeed(mycar car){
	fmt.Println("outside getspeed func")
	fmt.Println(mycar.Speed)
	mycar.Speed=80
	fmt.Println(mycar.Speed)
}


