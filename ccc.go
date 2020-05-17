package main

import (
	"fmt"
)

func main() {

	table := [5]int{}
	player := 55
	var ver int
	fmt.Println("please enter ver value:")
	fmt.Scanln(&ver)

	table[ver] = player

	fmt.Println(table[0])
	fmt.Println(table[1])
	fmt.Println(table[2])
}
