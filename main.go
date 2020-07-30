package main

import "fmt"

func main(){
	i := 1
	j := 99
	fmt.Println(&i,&j)
	i, j = j ,i
	fmt.Println(i, j)
	fmt.Println(&i,&j)
	//scrapper.Scrape("scala")
}

