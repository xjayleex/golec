package main

import(
	"./accounts"
	"fmt"
	"log"
)
//import "github.com/xjayleex/golec/accounts"
func main(){
	account := accounts.NewAccount("jaehyun")
	fmt.Println(account)
	account.Deposit(30)
	fmt.Println(account.Balance())
	err := account.Withdraw(40)
	if err!= nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
