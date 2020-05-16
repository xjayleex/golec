package main
import (
	"./urlchecker"
)
func main() {
	urls := [] string {
		"http://www.airbnb.com/",
		"http://www.naver.com/",
	}
	urlchecker.DoCheck(urls)

}
