package main

import (
	"fmt"

	bank "github.com/asackan/study-golang/example2_bank/accounts"
)

func main() {
	account := bank.Accounts{Owner: "owner", Balance: 1000}
	fmt.Println(account)
}
