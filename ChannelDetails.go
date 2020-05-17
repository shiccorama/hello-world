package main
import ("fmt"
	    
)

func local(c chan int, n int){
	c <- n*5
}

func main(){
	global:=make(chan int)

	go local(global,4)
	go local(global,5)
	go local(global,6)
	go local(global,7)
	go local(global,8)

	x1:=<-global
	x2:=<-global
	x3:=<-global
	x4:=<-global
	x5:=<-global

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)


}