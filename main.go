package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"jay" , "kay"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

func sexyCount(person string){
	for i:=0; i < 10 ; i++{
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}