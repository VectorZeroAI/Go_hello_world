package main
// Importing everything
import (
	"fmt"
	"math/rand"
)
type hello_world struct{
	words[]string
}
// Comment
func main() {
	tmp := hello_world{
		words:[]string{"hello", "world"},
	}
	messy_word := string(rand.Int())
	for len(tmp.words) <= 2{
		fmt.Println(tmp.words)
		tmp.words = append(tmp.words, messy_word)
	}
	if !(len(tmp.words) <= 2){
		fmt.Println(tmp.words)
		fmt.Println("As you see, this is a mess. ")
		fmt.Println("As far as I understand, ideal data storage is a slice of instanses of data. ")
		fmt.Println("an Instanse may be a Struct, or a lot of stuff. ")
	}
}
