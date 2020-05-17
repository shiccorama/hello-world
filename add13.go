package main
import "fmt"

func main(){

var Names[10] int
Names[0]=12
Names[1]=16
Names[4]=18

var Numb=[7] int {0,1,3,5,7,9,11}

var Dumb=[4] int{2,4,6,8}

for i:= range Numb{
fmt.Println(Numb[i])
}

for j:= range Names{
fmt.Println(Names[j])
}

for k:= range Dumb{
fmt.Println(Dumb[k])
}

fmt.Println(Dumb[0:4])




}