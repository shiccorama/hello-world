package main
import "fmt"
func main(){
for i:=0;i<10;i++{
fmt.Println(i)
if i==7{
	fmt.Println(" this is enough")
	break
}
}
var rate int=1
switch rate {
	case 0:
	fmt.Println("bad")
	fmt.Println("no star")
	case 1:
	fmt.Println("good")
	fmt.Println("one star")
	case 2:
	fmt.Println("very good")
	fmt.Println("two stars")
	case 3:
	fmt.Println("Excellent")
	fmt.Println("three stars")
	default:
	fmt.Println("out of the range")
	
}
}

