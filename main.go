package main

import(
	"./accounts"
	"fmt"
)
//import "github.com/xjayleex/golec/accounts"
func main(){
	account := accounts.NewAccount("jaehyun")
	fmt.Println(account)
	account.Deposit(30)
	fmt.Println(account.Balance())
}
