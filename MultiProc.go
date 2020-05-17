package main
import ("fmt"
        "sync")
		
		var wg sync.WaitGroup

func main(){
	

	

wg.Add(1)
go show("hello")
wg.Add(1)
go show("world")

wg.Wait()


}


func show(s string){

	for i:=0;i<3;i++{
		fmt.Println(s)
	}
		wg.Done()
}

