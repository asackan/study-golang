package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("* Can't withdraw. You don't have that much\n")

// NewAccount creats Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("Try Deposit", amount)
	a.balance += amount
	fmt.Println("* clear\n")
}

// return Balance of your Account
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw x ammount from your Account
func (a *Account) Withdraw(amount int) error {
	fmt.Println("Try Withdraw", amount)

	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	fmt.Println("* clear\n")
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	fmt.Print("Change Account Owner : ", a.owner, " to ")
	a.owner = newOwner
	fmt.Println(a.owner, "\n")
}

// return Owner name
func (a *Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. Has:", a.Balance())
}
