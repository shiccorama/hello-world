package main
import"fmt"
func main(){
	var Names[5] string
	Names[0]="Mohamed"
	Names[1]="Ahmed"
	Names[2]="Elhasher"
	
	var Numb[6] int
    Numb[0]=6	
	Numb[1]=7
	Numb[2]=8
	
	for i:=0;i<3;i++ {
	fmt.Println(Names[i])	
	}
	for j:=0;j<4;j++{
	fmt.Println(Numb[j])
	}
	
	
	
	fmt.Println(Names)
	fmt.Println(Numb)
	
	
	
	
	
	
}