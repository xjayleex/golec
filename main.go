package main

import "fmt"
import combi "./combination"

func main(){
	cb := combi.New()
	cb.SetParams(5,3)
	li := cb.GetLists()
	fmt.Println(li.Len())
	for li.Len() > 0 {
		elem := li.Front()
		fmt.Println(elem.Value)
		li.Remove(elem)
	}
}

