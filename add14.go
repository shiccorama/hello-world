package main
import "fmt"

var z=300

func main(){
var x=5
var y=10
fmt.Println(x,y,z)

TestGlobal()

fmt.Println(x,y)
}

func TestGlobal(){
var x=50
fmt.Println(x,z)

}
