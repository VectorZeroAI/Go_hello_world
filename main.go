/*
multiline comment
*/
package main

// Importing everything
import (
	"fmt"
)

type hello_world struct{
	words[]string
}
// Comment
func main() {
	tmp := hello_world{
		words:[]string{"hello", "world"},
	}
	for len(tmp.words) <= 2{
		fmt.Println(tmp.words)
	}
}
