package main
import "fmt"
func main(){
	
	var names[4] int
	names[0]=8
	names[1]=7
	names[2]=4
	names[3]=9
	
	for i:= range names{
		fmt.Println(names[i])
	}
}