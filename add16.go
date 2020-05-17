package main
import "fmt"

func main(){

var a=[] int{0,1,3,5,7}
b:=[] string{"*","#","$","&"}

fmt.Println(b[0:4])

for i:=range a{

fmt.Println(i)
fmt.Println(a[0:5])


}

x:=5
var y int
var p *int

y=x

p=&x
y=8
var z float64
z=float64(x)/float64(y)

fmt.Println(x)
fmt.Println(y)

*p=77
fmt.Println(x)
fmt.Println(z)


grades:=make(map[float64]string)
grades[6.44]="good"
grades[7.88]="very good"
grades[9.11]="excellent"

fmt.Println(grades)
fmt.Println(grades[9.11])


}