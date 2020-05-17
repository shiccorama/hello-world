package main
import "fmt"

func main(){


playerName := "hassan"
playerName = "hoksha"

if playerName=="ahmed"{
	fmt.Println("ahmed")
}else if playerName=="hoksha"{  /* note that you should write it that way }else  */
	fmt.Println("hoksha")
}else if playerName=="hamed"{/* note that if the 1st else if = 2nd else if, the program will execute the 1st and then terminate  */
	fmt.Println("hamed")
}else{                         /* note that else is written inside the flipped parenthesis }else{  */
	fmt.Println("no body is here")
}


if playerName!="islam"{         /* note : != means anything except "islam" */
	fmt.Println("I do not know who you are?")
}



}