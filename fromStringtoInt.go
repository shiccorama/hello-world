/*M
Programming / Golang
How to convert a string to an int type in Golang

 

 

 

How to convert a string to an int type in Golang - images/logos/golang.jpg

 

Here is a simple snippet how to convert a string into an int type in Golang. If you need help how to install Golang check the references links.
Code

To convert a string to an int you can use one of the following methods:

    strconv.Atoi: Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
    strconv.ParseInt: ParseInt interprets a string s in the given base (2 to 36) and returns the corresponding value i.*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := "1234"

	/** converting the str1 variable into an int using Atoi method */
	i1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(i1)
	}

	str2 := "5678"
	/** converting the str2 variable into an int using ParseInt method */
	i2, err := strconv.ParseInt(str2, 10, 64)
	if err == nil {
		fmt.Println(i2)
	}

/*the default values of variables without assigning any values are like this*/
	var x string
	var y int
	var z float64
	var zz float32
	var b bool

	fmt.Println(x,y,z,zz,b)
	

	message := make(chan string)
	go func(){
		message <- "hello my local watchers"
	}()
	msg := <-message
	fmt.Println(msg)
}
