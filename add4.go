package main
import "fmt"
func Head_Str(s1,s2 string) (string){
return s1+s2
}
func main(){
s1:= "Hi"
s2:= "Arabi"
fmt.Println(Head_Str(s1,s2))
}