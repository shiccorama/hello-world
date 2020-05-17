package main
import (
	"fmt"
	"time"
)

func main(){

age:=55

switch{
case age>18:
	fmt.Println("good")

	fallthrough /* this is a special condition of switch where you can neglect var after switch word and write the cases directly
	              without prior notification. Also, there is something called "fallthrough", this means that if you meet the 1st condition
	              and print it Don't stop, continue the program, so , if you run this program you will get (good and not bad)
	              but you will not get (terminate) unless you write it just before case age>52 */
case age>50:
	fmt.Println("not bad")
case age>52:
	fmt.Println("terminate")
}

label:= 8     /*this is the normal switch block of code BOC */

switch label {
case 4:
	fmt.Println("4 is not 5")
case 5:
	fmt.Println("5")
default:
	fmt.Println("hey big boy")


switch time.Now().Weekday(){            /* another solution is :*/
                                          
case time.Friday, time.Saturday:
	fmt.Println("it is a Week-end")
default:
	fmt.Println("it is a Weekday")
}

  if time.Now().Weekday()==time.Friday || time.Now().Weekday()==time.Saturday{
    fmt.Println("this is like special switch")
}

}



}