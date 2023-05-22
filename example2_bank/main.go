package main

import (
	"fmt"

	"github.com/asackan/study-golang/example2_bank/accounts"
)

func main() {
	account := accounts.NewAccount("juwon")

	account.Deposit(10)
	fmt.Println(account.Owner(), "have", account.Balance(), "\n")

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

	account.ChangeOwner("mom")
	fmt.Println(account)
}
