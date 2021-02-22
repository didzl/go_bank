package main

import (
	"fmt"
	"hello/bankTest/accounts"
)


func main() {
	account:= accounts.NewAccount("Han")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}