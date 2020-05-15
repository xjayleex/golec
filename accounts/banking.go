package accounts

//Account struct
type Account struct {
	owner string
	balance int
}

// Creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}
// Method
// Deposit x amount on an account
// a 는 Receiver
func (a Account) Deposit(amount int){
	a.balance += amount
}