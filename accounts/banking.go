package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner string
	balance int
}

var errorNoMoney = errors.New("Can't withdraw")
// Creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}
// Method
// Deposit x amount on an account
// a 는 Receiver
func (a *Account) Deposit(amount int){
	a.balance += amount
}

func (a Account) Balance() int{
	return a.balance
}

// Withdraw from account
func (a *Account) Withdraw(amount int) error{
	if a.balance < amount {
		return errorNoMoney
	}
	a.balance -= amount
	return nil
}

//Change owner
func (a *Account) ChangeOwner(newOwner string){
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string { // equiv toString()
	return fmt.Sprint(a.Owner(), "'s account.\nHas:",
		a.balance)
}