package main
import ("fmt"
	    
)

func local (c chan int, n int) {
	c <- n*5
}

func main(){


yy:=make(chan int)

go local(yy,1)
go local(yy,2)
go local(yy,3)
go local(yy,4)
go local(yy,5)

	y1:=<-yy
	y2:=<-yy
	y3:=<-yy
	y4:=<-yy
	y5:=<-yy

	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
	fmt.Println(y5)

}