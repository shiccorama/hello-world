package main
import (
	"fmt"
	"net/http"
)

	func main() {

		http.HandleFunc("/site/", MyServe)
		http.ListenAndServe(":8080", nil)
	}
	    
func MyServe(w http.ResponseWriter, hur *http.Request) {
	fmt.Fprintln(w, "this is my site")
	fmt.Fprintf(w, "this is %v my site","88")
}
