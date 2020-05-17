package main
import "fmt"

func main(){

x:=5
y:=8


var p *int
p=&x

y=3
*p=78

fmt.Println(x,y)

var grades=make(map[float64]int)
grades[6.8]=77
grades[4.5]=97
grades[8.2]=83

fmt.Println(grades[8.2])
fmt.Println(grades)


}