package main

// Desired output:
// this.menu.popup.menuitem[0].onclick = "CreateNewDoc()"
// this.menu.popup.menuitem[0].value = "New"
// this.menu.popup.menuitem[1].onclick = "OpenDoc()"
// this.menu.popup.menuitem[1].value = "Open"
// this.menu.popup.menuitem[2].onclick = "CloseDoc()"
// this.menu.popup.menuitem[2].value = "Close"
// this.menu.id = "file"
// this.menu.value = "File"

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func dump(prefix string, i interface{}) {
	switch ob := i.(type) {
	case map[string]interface{}:
		for k, v := range ob {
			dump(prefix+"."+k, v)
		}
	case []interface{}:
		for i, v := range ob {
			dump(prefix+"["+strconv.Itoa(i)+"]", v)
		}
	case string:
		fmt.Printf("%s = %q\n", prefix, ob)
	default:
		fmt.Printf("Unhandled: %T\n", i)
	}
}

func main() {
	var j = []byte(`
{"menu": {
  "id": "file",
  "value": "File",
  "popup": {
    "menuitem": [
      {"value": "New", "onclick": "CreateNewDoc()"},
      {"value": "Open", "onclick": "OpenDoc()"},
      {"value": "Close", "onclick": "CloseDoc()"}
    ]
  }
}}
`)

	var myOb interface{}
	err := json.Unmarshal(j, &myOb)
	if err != nil {
		panic(err)
	}
	dump("this", myOb)
}