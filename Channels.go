package main
import "fmt"

func main(){

	val1:=make(map[float64]string)
    val1[0.65]="good"
	val1[0.75]="very good"
	val1[0.85]="excellent"
    fmt.Println(val1)

	
    val2:=make(chan int)

	go foo(val2,700)
	go foo(val2,8)
	go foo(val2,9)
	
	x:=<-val2
	y:=<-val2
	z:=<-val2

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)


}

func foo(c chan int, n int){

	c <- n*4

}