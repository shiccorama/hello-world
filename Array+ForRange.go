package main
import "fmt"


func main(){

var a=[5] int{2,4,6,8,10}

for i:=range a{
fmt.Println(i)
}
fmt.Println(a[2:4])
fmt.Println(a[0:5])

fmt.Println("out of Discount")
}
