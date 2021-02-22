package accounts

import "errors"

//account banking
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Cant withdraw")

//NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance :0}
	return &account
}

//deposit method
func (abc *Account) Deposit(amount int) {
	abc.balance += amount
}

//balance check method
func (a Account) Balance() int {
	return a.balance
}

// withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount{
		return errNoMoney
	}
	a.balance -= amount
	return nil
}