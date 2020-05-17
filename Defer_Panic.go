package main
import ("fmt")
		
	

func main(){
	
	
	x:=10
	y:=0
	z:=0;
	
	defer fmt.Println(x,y,z)
	
	
	if(y==0){
		panic("divided by zero is not allowed")
	}else{
		z=x/y
	}
    
	
	fmt.Println(z)
	
	}

