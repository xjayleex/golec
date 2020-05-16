package urlchecker

import (
	"errors"
	"fmt"
	"net/http"

)
type ResultMap map[string]int

var (
	errRequestFailed = errors.New("Req failed.")
)


func hitURL(url string) int {


	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		return resp.StatusCode
	}
	return 200
}

func DoCheck(urls [] string){
	resultMap := ResultMap{}
	for _, url := range urls{
		code := hitURL(url)
		resultMap[url] = code
	}
	fmt.Println(resultMap)
}
