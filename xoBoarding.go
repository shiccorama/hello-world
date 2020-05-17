package main

import (
	"fmt"
)

func main() {

	xoBoard := [3][3]string{}
	player := "x"
	fmt.Println("Welcome to xo game, 1st player will be:", player)

	var row int
	fmt.Println("enter row number")
	fmt.Scanln(&row)
	var col int
	fmt.Println("enter column number")
	fmt.Scanln(&col)
	fmt.Println("row =", row, "column=", col)

	xoBoard[row][col] = player

}
