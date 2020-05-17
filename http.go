package main
import (
	"fmt"
	"net/http"
)

	func main() {

		http.HandleFunc("/hhhh/", MySun)
		http.ListenAndServe(":8080", nil)
	}
	    
func MySun(m http.ResponseWriter, mon *http.Request) {
	fmt.Fprintln(m, "this is my door")
	fmt.Fprintf(m, "this is %v my site","99%")
}

/* note that : ("/http/"+ m + mon) are variable can be renamed as you wish
and also the function (MySun) can be renamed , but (HandleFunc,ListenAndServe,http.ResponseWriter,*http.Request,Fprintln,Fprintf)
cannot be renamed because they are functions inside "fmt"and"net/http".
check http1.go as I renamed all */
