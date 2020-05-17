package main
import "fmt"

func main() {


	xoBoard:=[3][3]string{

	[3]string{"x","-","o"},
	[3]string{"o","x","-"},
	[3]string{"o","-","x"},
}

fmt.Println(xoBoard)

fmt.Println(xoBoard[0])
fmt.Println(xoBoard[1])
fmt.Println(xoBoard[2])

for i:=0;i<3;i++{
	for j:=0;j<3;j++{
		fmt.Printf("x ")
	}
	fmt.Println("")
}






}
