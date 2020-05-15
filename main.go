package main
import (
	"./dict"
	"fmt"
)
func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(definition)
	}
}
