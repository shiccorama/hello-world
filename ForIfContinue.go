package main

import "fmt"

/* exception loop*/
func dal() {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("next is breakpoint loop")
}
func main() {
	dal()

	/* breakpoint: loop*/

breakpoint:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break breakpoint
			}
			fmt.Println(x, y)

		}

	}

	/* boolean var loop*/

	fmt.Println("next is boolean example")
	b := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				b = true
				break
			}
			fmt.Println(x, y)
		}

		if b == true {
			break
		}

	}

}
