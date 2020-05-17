package main
import "fmt"
func main(){
	message("Omayma")
	message("Ibrahim")
	message("Ahmed")
	var x=5
	var y=7
	var z=Add(x,y)
	var d=3
	
	fmt.Println(z)
	fmt.Println(Add_Minus(x,y,d))
}
func message(Name string){
	fmt.Println("hello",Name)
}
func Add(x int, y int)int{
	return x+y
}
func Add_Minus(x int,y int,d int)(int,int,int){
	return x-y,x+y,y-x
	
}