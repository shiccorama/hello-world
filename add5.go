package main
import "fmt"
func Add(x int,y float32) (float32) {
return float32(x)+y
}
func main(){
n1:= 7.7
n2:= int(n1)
n3:= float32(n1)
fmt.Println(Add(n2,n3))
}