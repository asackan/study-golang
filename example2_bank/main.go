package main

import (
	"fmt"

	bank "github.com/asackan/study-golang/example2_bank/accounts"
)

func main() {
	account := bank.NewAccount()
	fmt.Println(account)
}
