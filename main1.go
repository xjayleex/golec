package main
import (
	"./dict"
	"fmt"
)
func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
		dictionary.Add("second","secondWord")
	} else{
	}
	definition, err = dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err = dictionary.Update("second","updated")
}
