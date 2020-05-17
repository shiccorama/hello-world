package main
import (
	"fmt"
	"net/http"
)

	func main() {

		http.HandleFunc("/site/", MyServe)
		http.ListenAndServe(":8080", nil)
	}
	    
func MyServe(w http.ResponseWriter, *http.Request) {
	fmt.Fprintln(w, "this is my site")
	fmt.Fprintf(w, "this is %s my site","54")


}
